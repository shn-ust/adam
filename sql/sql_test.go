package sql

import (
	"UST-FireOps/adam/utils"
	"net"
	"testing"
)

func TestInsertPacket(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		sourcePort = uint16(8080)
		destIP     = net.ParseIP("127.0.0.1")
		destPort   = uint16(5432)
	)
	packet := utils.CreatePacket(sourceIP, destIP, sourcePort, destPort)

	if ok := InsertPacket(packet); !ok {
		t.Errorf("Failed to insert packet data to table!")
	}

}
