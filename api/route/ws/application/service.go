package application

import (
	"net/http"
	"time"

	"auctionsystem/api/route/ws/domain"
	"auctionsystem/api/route/ws/infra/mq"
	"auctionsystem/pkg/config"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type AuctionHubService struct {
	hub      *domain.AuctionHub
	upgrader websocket.Upgrader
}

func NewAuctionHubService(redis *redis.Client) *AuctionHubService {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  config.WSReadBufferSize,
		WriteBufferSize: config.WSWriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	mq := mq.NewRedisRepository(redis)
	return &AuctionHubService{
		hub:      domain.NewAuctionHub(mq),
		upgrader: upgrader,
	}
}

func (s *AuctionHubService) Upgrade(cmd UpgradeCommand) (*websocket.Conn, error) {
	return s.upgrader.Upgrade(cmd.Writer, cmd.Request, nil)
}

func (s *AuctionHubService) RegisterClient(auctionID uint, conn *websocket.Conn) *Client {
	client := domain.NewAuctionClient(conn, auctionID, s.hub.Unregister)
	s.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return &Client{
		AuctionID: auctionID,
		Conn:      conn,
		send:      client.Send,
	}
}

func (s *AuctionHubService) SendFirstBid(cmd SendBidCommand) {
	cmd.Client.send <- domain.AuctionMessage{
		AuctionID: cmd.AuctionID,
		BidInfo: domain.BidInfo{
			Price:     cmd.Price,
			Timestamp: uint(time.Now().Unix()),
		},
	}
}

func (s *AuctionHubService) RunHub() {
	s.hub.Run()
}
