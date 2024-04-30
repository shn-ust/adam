package sql

import (
	"time"

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

func InsertPacket(packet parse.ParsedPacket, db *gorm.DB) bool {
	db.AutoMigrate(&Flow{})

	res := db.Create(&Flow{
		Timestamp:  packet.TimeStamp,
		SourceIP:   packet.SourceIP.String(),
		SourcePort: uint16(packet.SourcePort),
		DestIP:     packet.DestinationIP.String(),
		DestPort:   uint16(packet.DestinationPort),
	})

	return res.Error == nil
}
