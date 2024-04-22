package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	const snapLen = 262144

	handle, err := pcap.OpenLive("eth0", snapLen, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	// Filter out SSH connections
	if err := handle.SetBPFFilter("port not 22"); err != nil {
		panic(err)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		ParsePacket(packet)
	}
}
