package parse

import (
	"net"
	"reflect"
	"testing"

	"github.com/google/gopacket/layers"

	"UST-FireOps/adam/utils"
)

func TestParsePacket(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		destIP     = net.ParseIP("127.0.0.1")
		sourcePort = uint16(8080)
		destPort   = uint16(5432)
	)

	packet := utils.CreatePacket(sourceIP, destIP, sourcePort, destPort)

	got := ParsePacket(packet)
	want := ParsedPacket{
		SrcIP:    sourceIP,
		SrcPort:  layers.TCPPort(sourcePort),
		DestIP:   destIP,
		DestPort: layers.TCPPort(destPort),
	}

	if !reflect.DeepEqual(got.SrcPort, want.SrcPort) {
		t.Errorf("SourcePort mismatch: got %+v, want %+v", got.SrcPort, want.SrcPort)
	}

	if !reflect.DeepEqual(got.DestPort, want.DestPort) {
		t.Errorf("DestinationPort mismatch: got %+v, want %+v", got.DestPort, want.DestPort)
	}

	if got.SrcIP.String() != want.SrcIP.String() {
		t.Errorf("SourceIP mismatch: got %+v, want %+v", got.SrcIP, want.SrcIP)
	}

	if got.DestIP.String() != want.DestIP.String() {
		t.Errorf("DestinationIP mismatch: got %+v, want %+v", got.DestIP, want.DestIP)
	}
}
