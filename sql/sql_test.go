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
		panic("Failed to connect to database")
	}

	parsedPacket := parse.ParsePacket(packet)

	var mu sync.Mutex

	if ok := InsertPacket(parsedPacket, db, &mu); !ok {
		t.Errorf("Failed to insert packet data to table!")
	}

}
