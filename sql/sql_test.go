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

func TestInsertPacketsInBatch(t *testing.T) {
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

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	db.AutoMigrate(&PacketDetail{})

	parsedPacket := parse.ParsePacket(packet)
	parsedPacket2 := parse.ParsePacket(packet2)

	var mu sync.Mutex

	packets := []*parse.ParsedPacket{parsedPacket, parsedPacket2}
	if err := InsertPacketsInBatch(db, &mu, packets); err != nil {
		t.Errorf("Error inserting data in batches: %v", err)
	}

	var count int64
	if err := db.Model(&PacketDetail{}).Count(&count).Error; err != nil {
		t.Fatalf("Error counting records: %v", err)
	}

	if count != 2 {
		t.Errorf("Count mismatch, got: %v want: 2", count)
	}

	var packetDetail1 PacketDetail
	if err := db.Where("src_ip = ? AND src_port = ? AND dest_ip = ? AND dest_port = ?",
		sourceIP.String(), sourcePort, destIP.String(), destPort).First(&packetDetail1).Error; err != nil {
		t.Errorf("Expected packet1 to be in database, but got error: %v", err)
	}

}
