package domain

import (
	"auctionsystem/pkg/config"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type AuctionClient struct {
	conn *websocket.Conn

	Unregister chan<- *AuctionClient
	Send       chan AuctionMessage

	auctionID uint
}

func NewAuctionClient(conn *websocket.Conn, auctionID uint, unregister chan<- *AuctionClient) *AuctionClient {
	return &AuctionClient{
		conn:       conn,
		auctionID:  auctionID,
		Unregister: unregister,
		Send:       make(chan AuctionMessage),
	}
}

func (c *AuctionClient) Close() {
	c.Unregister <- c
	close(c.Send)
	c.conn.Close()
}

func (c *AuctionClient) WritePump() {
	ticker := time.NewTicker(config.WSWriteTick * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(config.WSWriteDeadline * time.Second))
			if !ok {
				return
			}

			writer, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				return
			}

			if _, err = writer.Write(jsonMsg); err != nil {
				return
			}

			if err := writer.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(config.WSWriteDeadline * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *AuctionClient) ReadPump() {
	defer c.Close()

	// 防止超大消息导致内容膨胀
	c.conn.SetReadLimit(config.WSReadBufferSize)
	// 超时处理
	_ = c.conn.SetReadDeadline(time.Now().Add(config.WSReadDeadline * time.Second))
	c.conn.SetPongHandler(func(appData string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(config.WSReadDeadline * time.Second))
		return nil
	})

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
