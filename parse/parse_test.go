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
