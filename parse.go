package main

import (
	// "fmt"

	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ParsedPacket struct {
	TimeStamp       time.Time
	SourceIP        net.IP
	SourcePort      layers.TCPPort
	DestinationIP   net.IP
	DestinationPort layers.TCPPort
}

// Used to extract the necessary informations from a packet
func ParsePacket(packet gopacket.Packet) ParsedPacket {
	// timeStamp := packet.Metadata().Timestamp
	fmt.Println(packet)

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		fmt.Println("No ip layer!")
		fmt.Println(packet)
	}

	return ParsedPacket{}
}
