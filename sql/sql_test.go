package sql

import (
	"UST-FireOps/adam/parse"
	"UST-FireOps/adam/utils"
	"net"
	"sync"
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
		t.Error("Failed to connect to database")
	}

	parsedPacket := parse.ParsePacket(packet)

	var mu sync.Mutex

	if ok := InsertPacket(parsedPacket, db, &mu); !ok {
		t.Error("Failed to insert packet data to table!")
	}
}

func TestInsertPacketInBatch(t *testing.T) {
	var (
		sourceIP   = net.ParseIP("127.0.0.1")
		sourcePort = uint16(8080)
		destIP     = net.ParseIP("127.0.0.1")
		destPort   = uint16(5432)

		sourceIP2   = net.ParseIP("127.0.0.2")
		sourcePort2 = uint16(8080)
		destIP2     = net.ParseIP("127.0.0.2")
		destPort2   = uint16(5432)
	)
	packet := utils.CreatePacket(sourceIP, destIP, sourcePort, destPort)
	packet2 := utils.CreatePacket(sourceIP2, destIP2, sourcePort2, destPort2)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		t.Error("Failed to connect to database")
	}

	parsedPacket := parse.ParsePacket(packet)
	parsedPacket2 := parse.ParsePacket(packet2)

	var mu sync.Mutex

	packets := []*parse.ParsedPacket{parsedPacket, parsedPacket2}
	if err := InsertPacketInBatch(db, &mu, packets); err != nil {
		t.Errorf("Error inserting data in batches: %v", err)
	}

}
