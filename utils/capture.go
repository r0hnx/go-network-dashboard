package utils

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func StartPacketCapture(device string) {
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if packet.TransportLayer().LayerType().String() == "TCP" {
			fmt.Println(packet)
		}
		// Process packet here
		// fmt.Println(packet)
	}
}
