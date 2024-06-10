package sql

import (
	"sync"
	"time"

	"gorm.io/gorm"

	"UST-FireOps/adam/parse"
)

type PacketDetail struct {
	ID        uint `gorm:"primarykey"`
	Timestamp time.Time
	SrcIP     string
	SrcPort   uint16
	DestIP    string
	DestPort  uint16
}

// convertParsedPacket is used to create an array of "PacketDetail"
// from an array of "ParsedPacket"
func convertParsedPacket(packets []*parse.ParsedPacket) []PacketDetail {
	var details []PacketDetail

	for _, packet := range packets {
		details = append(details, PacketDetail{
			Timestamp: packet.TimeStamp,
			SrcIP:     packet.SrcIP.String(),
			SrcPort:   uint16(packet.SrcPort),
			DestIP:    packet.DestIP.String(),
			DestPort:  uint16(packet.DestPort),
		})
	}

	return details
}

// InsertPacketInBatch inserts an array of packets into the database
func InsertPacketsInBatch(db *gorm.DB, mutex *sync.Mutex, packets []*parse.ParsedPacket) error {
	mutex.Lock()
	defer mutex.Unlock()

	const batchSize = 256

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(convertParsedPacket(packets), batchSize).Error; err != nil {
			return err
		}

		return nil
	})
}
