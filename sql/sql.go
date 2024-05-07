package sql

import (
	"sync"
	"time"

	"gorm.io/gorm"

	"UST-FireOps/adam/parse"
)

type PacketDetail struct {
	gorm.Model
	Timestamp time.Time
	SrcIP     string
	SrcPort   uint16
	DestIP    string
	DestPort  uint16
}

func InsertPacket(packet parse.ParsedPacket, db *gorm.DB, mutex *sync.Mutex) bool {
	mutex.Lock()
	defer mutex.Unlock()

	db.AutoMigrate(&PacketDetail{})

	res := db.Create(&PacketDetail{
		Timestamp: packet.TimeStamp,
		SrcIP:     packet.SrcIP.String(),
		SrcPort:   uint16(packet.SrcPort),
		DestIP:    packet.DestIP.String(),
		DestPort:  uint16(packet.DestPort),
	})

	return res.Error == nil
}
