package main

import (
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
	FIN             bool
	SYN             bool
	RST             bool
	PSH             bool
	ACK             bool
	URG             bool
	ECE             bool
	CWR             bool
}

var (
	ethernet layers.Ethernet
	ipv4     layers.IPv4
	tcp      layers.TCP
	payload  gopacket.Payload
)

// Used to extract the necessary informations from a packet
func ParsePacket(packet gopacket.Packet) ParsedPacket {
	timeStamp := packet.Metadata().Timestamp

	var (
		sourceIP        net.IP
		sourcePort      layers.TCPPort
		destionationIP  net.IP
		destinationPort layers.TCPPort
		fin             bool
		syn             bool
		rst             bool
		psh             bool
		ack             bool
		urg             bool
		ece             bool
		cwr             bool
	)

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ethernet, &ipv4, &tcp, &payload)
	decoded := []gopacket.LayerType{}

	err := parser.DecodeLayers(packet.Data(), &decoded)

	if err != nil {
		panic(err)
	}

	for _, layerType := range decoded {
		if layerType == layers.LayerTypeIPv4 {
			sourceIP = ipv4.SrcIP
			destionationIP = ipv4.DstIP
		} else if layerType == layers.LayerTypeTCP {
			sourcePort = tcp.SrcPort
			destinationPort = tcp.DstPort
			fin = tcp.FIN
			syn = tcp.SYN
			rst = tcp.RST
			psh = tcp.PSH
			ack = tcp.ACK
			urg = tcp.URG
			ece = tcp.ECE
			cwr = tcp.CWR
		}
	}

	return ParsedPacket{
		TimeStamp:       timeStamp,
		SourceIP:        sourceIP,
		SourcePort:      sourcePort,
		DestinationIP:   destionationIP,
		DestinationPort: destinationPort,
		FIN:             fin,
		SYN:             syn,
		RST:             rst,
		PSH:             psh,
		ACK:             ack,
		URG:             urg,
		ECE:             ece,
		CWR:             cwr,
	}
}
