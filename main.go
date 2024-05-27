package main

import (
	"log"
	"net/http"
	"nmap-scanner/handler"
	"nmap-scanner/utils"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/nmap", handler.HandleNmapScan)
	mux.HandleFunc("/api/traceroute", handler.TracerouteHandler)
	mux.HandleFunc("/ws", utils.WsHandler)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", handler.EnableCors(mux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
