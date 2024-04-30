package main

import (
	"net"
	"reflect"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Helper function to create packet
func createPacket(sourceIP, destIP net.IP, sourcePort, destPort uint16) gopacket.Packet {
	ethernetLayer := layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		DstMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		EthernetType: layers.EthernetTypeIPv4,
	}

	ipLayer := layers.IPv4{
		SrcIP:    sourceIP,
		DstIP:    destIP,
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(sourcePort),
		DstPort: layers.TCPPort(destPort),
	}

	tcpLayer.SetNetworkLayerForChecksum(&ipLayer)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	if err := gopacket.SerializeLayers(buf, opts, &ethernetLayer, &ipLayer, &tcpLayer); err != nil {
		panic(err)
	}

	packet := gopacket.NewPacket(buf.Bytes(), layers.LayerTypeEthernet, gopacket.Default)

	return packet
}

func TestParsePacket(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		destIP     = net.ParseIP("127.0.0.1")
		sourcePort = uint16(8080)
		destPort   = uint16(5432)
	)

	packet := createPacket(sourceIP, destIP, sourcePort, destPort)

	got := ParsePacket(packet)
	want := ParsedPacket{
		SourceIP:        sourceIP,
		SourcePort:      layers.TCPPort(sourcePort),
		DestinationIP:   destIP,
		DestinationPort: layers.TCPPort(destPort),
	}

	if !reflect.DeepEqual(got.SourcePort, want.SourcePort) {
		t.Errorf("SourcePort mismatch: got %+v, want %+v", got.SourcePort, want.SourcePort)
	}

	if !reflect.DeepEqual(got.DestinationPort, want.DestinationPort) {
		t.Errorf("DestinationPort mismatch: got %+v, want %+v", got.DestinationPort, want.DestinationPort)
	}

	if got.SourceIP.String() != want.SourceIP.String() {
		t.Errorf("SourceIP mismatch: got %+v, want %+v", got.SourceIP, want.SourceIP)
	}

	if got.DestinationIP.String() != want.DestinationIP.String() {
		t.Errorf("DestinationIP mismatch: got %+v, want %+v", got.DestinationIP, want.DestinationIP)
	}
}
