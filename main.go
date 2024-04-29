package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	const snapLen = 262144

	networkInterface := "lo"
	handle, err := pcap.OpenLive(networkInterface, snapLen, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	// if err := handle.SetBPFFilter("port 5000 or port 5001"); err != nil {
	// 	panic(err)
	// }

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Printf("%+v\n", ParsePacket(packet))
	}
}
