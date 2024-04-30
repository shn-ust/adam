package sql

import (
	"time"

	"github.com/google/gopacket"
	"gorm.io/gorm"

	"UST-FireOps/adam/parse"
)

type Flow struct {
	gorm.Model
	Timestamp  time.Time
	SourceIP   string
	SourcePort uint16
	DestIP     string
	DestPort   uint16
}

func InsertPacket(packet gopacket.Packet, db *gorm.DB) bool {
	db.AutoMigrate(&Flow{})

	parsedPacket := parse.ParsePacket(packet)

	res := db.Create(&Flow{
		Timestamp:  parsedPacket.TimeStamp,
		SourceIP:   parsedPacket.SourceIP.String(),
		SourcePort: uint16(parsedPacket.SourcePort),
		DestIP:     parsedPacket.DestinationIP.String(),
		DestPort:   uint16(parsedPacket.DestinationPort),
	})

	return res.Error == nil
}
