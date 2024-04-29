package main

import (
	"net"
	"reflect"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestParsePacket(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		destIP     = net.ParseIP("127.0.0.1")
		sourcePort = 8080
		destPort   = 5432
	)

	ipLayer := &layers.IPv4{
		SrcIP: sourceIP,
		DstIP: destIP,
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(sourcePort),
		DstPort: layers.TCPPort(destPort),
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	if err := gopacket.SerializeLayers(buf, opts, ipLayer, tcpLayer); err != nil {
		panic(err)
	}

	packet := gopacket.NewPacket(buf.Bytes(), layers.LayerTypeIPv4, gopacket.Default)

	got := ParsePacket(packet)
	want := ParsedPacket{
		SourceIP:        sourceIP,
		SourcePort:      layers.TCPPort(sourcePort),
		DestinationIP:   destIP,
		DestinationPort: layers.TCPPort(destPort),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
