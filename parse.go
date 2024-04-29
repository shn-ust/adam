package main

import (
	"encoding/binary"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ParsedPacket struct {
	// TimeStamp       time.Time
	SourceIP        net.IP
	SourcePort      layers.TCPPort
	DestinationIP   net.IP
	DestinationPort layers.TCPPort
}

// Used to extract the necessary informations from a packet
func ParsePacket(packet gopacket.Packet) ParsedPacket {
	// timeStamp := packet.Metadata().Timestamp
	netFlow := packet.NetworkLayer().NetworkFlow()
	sourceIP, destIP := netFlow.Endpoints()

	transportFlow := packet.TransportLayer().TransportFlow()
	sourcePort, destPort := transportFlow.Endpoints()

	return ParsedPacket{
		// TimeStamp:       timeStamp,
		SourceIP:        sourceIP.Raw(),
		SourcePort:      layers.TCPPort(binary.BigEndian.Uint16(sourcePort.Raw())),
		DestinationIP:   destIP.Raw(),
		DestinationPort: layers.TCPPort(binary.BigEndian.Uint16(destPort.Raw())),
	}
}
