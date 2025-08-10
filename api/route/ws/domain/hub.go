package domain

import (
	"context"
	"encoding/json"
	"sync"
)

type AuctionHub struct {
	clients map[uint]map[*AuctionClient]struct{}

	Broadcast  chan AuctionMessage
	Register   chan *AuctionClient
	Unregister chan *AuctionClient

	repo     AuctionMessageRepository
	channels map[uint]MQChannel

	running bool

	mu sync.RWMutex
}

func NewAuctionHub(repo AuctionMessageRepository) *AuctionHub {
	return &AuctionHub{
		clients: make(map[uint]map[*AuctionClient]struct{}),

		Broadcast:  make(chan AuctionMessage),
		Register:   make(chan *AuctionClient),
		Unregister: make(chan *AuctionClient),

		repo:     repo,
		channels: make(map[uint]MQChannel),

		running: false,
		mu:      sync.RWMutex{},
	}
}

func (h *AuctionHub) readChannel(auctionID uint) {
	for msg := range h.channels[auctionID].Receive() {
		var message AuctionMessage
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			continue
		}
		h.Broadcast <- message
	}
}

func (h *AuctionHub) registerClient(client *AuctionClient) {
	h.mu.Lock()
	defer h.mu.Unlock()

	auctionID := client.auctionID
	if _, exists := h.clients[auctionID]; !exists {
		h.clients[auctionID] = make(map[*AuctionClient]struct{})
		h.channels[auctionID] = h.repo.Subscribe(context.Background(), auctionID)
		go h.readChannel(auctionID)
	}
	h.clients[auctionID][client] = struct{}{}
}

func (h *AuctionHub) unregisterClient(client *AuctionClient) {
	h.mu.Lock()
	defer h.mu.Unlock()

	auctionID := client.auctionID
	if _, exists := h.clients[auctionID]; !exists {
		return
	}
	delete(h.clients[auctionID], client)
	if len(h.clients[auctionID]) == 0 {
		delete(h.clients, auctionID)
		h.channels[auctionID].Cancel()
		delete(h.channels, auctionID)
	}
}

func (h *AuctionHub) broadcastMessage(messages AuctionMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for client := range h.clients[messages.AuctionID] {
		client.Send <- messages
	}
}

func (h *AuctionHub) Run() {
	h.mu.Lock()
	if h.running {
		h.mu.Unlock()
		return
	}
	h.running = true
	h.mu.Unlock()

	for {
		select {
		case client := <-h.Register:
			h.registerClient(client)
		case client := <-h.Unregister:
			h.unregisterClient(client)
		case messages := <-h.Broadcast:
			h.broadcastMessage(messages)
		}
	}
}
