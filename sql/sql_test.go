package sql

import (
	"UST-FireOps/adam/utils"
	"net"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInsertPacket(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		sourcePort = uint16(8080)
		destIP     = net.ParseIP("127.0.0.1")
		destPort   = uint16(5432)
	)
	packet := utils.CreatePacket(sourceIP, destIP, sourcePort, destPort)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	if ok := InsertPacket(packet, db); !ok {
		t.Errorf("Failed to insert packet data to table!")
	}

}
