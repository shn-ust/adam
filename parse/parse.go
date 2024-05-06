package parse

import (
	"encoding/binary"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ParsedPacket struct {
	TimeStamp time.Time
	SrcIP     net.IP
	SrcPort   layers.TCPPort
	DestIP    net.IP
	DestPort  layers.TCPPort
}

// Used to extract the necessary information from a packet
func ParsePacket(packet gopacket.Packet) ParsedPacket {
	timeStamp := packet.Metadata().Timestamp
	netFlow := packet.NetworkLayer().NetworkFlow()
	sourceIP, destIP := netFlow.Endpoints()

	transportFlow := packet.TransportLayer().TransportFlow()
	sourcePort, destPort := transportFlow.Endpoints()

	return ParsedPacket{
		TimeStamp: timeStamp,
		SrcIP:     sourceIP.Raw(),
		SrcPort:   layers.TCPPort(binary.BigEndian.Uint16(sourcePort.Raw())),
		DestIP:    destIP.Raw(),
		DestPort:  layers.TCPPort(binary.BigEndian.Uint16(destPort.Raw())),
	}
}
