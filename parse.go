package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)
import "fmt"

func ParsePacket(packet gopacket.Packet) {
	timeStamp := packet.Metadata().Timestamp
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)

	if ipLayer != nil && tcpLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		tcp, _ := tcpLayer.(*layers.TCP)

		fmt.Println(timeStamp, ": (", ip.SrcIP, tcp.SrcPort, ") -> (", ip.DstIP, tcp.DstPort, ")")
	}
}
