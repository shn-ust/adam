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

// ParsePacket is used to extract the necessary information from a packet
// The informations are:
// a. timestamp
// b. source ip and port
// c. destination ip and port
func ParsePacket(packet gopacket.Packet) *ParsedPacket {
	var (
		sourceIP   gopacket.Endpoint
		sourcePort gopacket.Endpoint
		destIP     gopacket.Endpoint
		destPort   gopacket.Endpoint
	)

	timeStamp := packet.Metadata().Timestamp
	networkLayer := packet.NetworkLayer()

	if networkLayer != nil {
		netFlow := networkLayer.NetworkFlow()
		sourceIP, destIP = netFlow.Endpoints()
	}

	transportLayer := packet.TransportLayer()
	if transportLayer != nil {
		transportFlow := transportLayer.TransportFlow()
		sourcePort, destPort = transportFlow.Endpoints()
	}

	if networkLayer != nil && transportLayer != nil {
		return &ParsedPacket{
			TimeStamp: timeStamp,
			SrcIP:     sourceIP.Raw(),
			SrcPort:   layers.TCPPort(binary.BigEndian.Uint16(sourcePort.Raw())),
			DestIP:    destIP.Raw(),
			DestPort:  layers.TCPPort(binary.BigEndian.Uint16(destPort.Raw())),
		}
	}

	return nil
}
