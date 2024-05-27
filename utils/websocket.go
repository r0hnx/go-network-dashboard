package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			packet := captureNetworkPacket()
			err := conn.WriteJSON(packet)
			if err != nil {
				log.Println("WriteJSON error:", err)
				return
			}
		}
	}
}

type NetworkPacket struct {
	Timestamp   time.Time `json:"timestamp"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Protocol    string    `json:"protocol"`
	Length      int       `json:"length"`
}

func captureNetworkPacket() NetworkPacket {
	// Simulate network packet capture
	return NetworkPacket{
		Timestamp:   time.Now(),
		Source:      "192.168.1.1",
		Destination: "192.168.1.100",
		Protocol:    "TCP",
		Length:      64,
	}
}
