package adm

import (
	"time"

	"UST-FireOps/adam/sql"

	"gorm.io/gorm"
)

// Flow is defined as a tuple containing
// (StartTime, EndTime, SourceIP, SourcePort, DestinationIP, DestinationPort)
// It represents when a request started and when it ended together with IP and Port details
type Flow struct {
	StartTime  time.Time
	EndTime    time.Time
	SourceIP   string
	SourcePort uint16
	DestIP     string
	DestPort   uint16
}

// This functions takes a list of packets (stored in SQLite inmemory database)
// and returns the 'Flow' of packets
func CreateFlow(db *gorm.DB) []Flow {
	var flows []Flow

	// Sort the packets in descending order based on timestamp
	var sortedPacketDetails []sql.PacketDetail
	db.Order("timestamp desc").Find(&sortedPacketDetails)

	// If the current packet has the greatest timestamp of all
	// Consider it as the last one
	// And find the first request with the same source ip, source port, destination ip, destination port
	for _, packet := range sortedPacketDetails {
		var firstOccurence sql.PacketDetail
		var count int64
		db.Model(&sql.PacketDetail{}).Where("source_ip = ?", packet.SourceIP).Where("source_port = ?", packet.SourcePort).Where("dest_ip = ?", packet.DestIP).Where("dest_port = ?", packet.DestPort).Where("timestamp > ?", packet.Timestamp).Count(&count)

		// '0', if there doesn't exists a packet having a greater timestamp than the current one
		if count == 0 {
			// Find the first occurence
			db.Where("source_ip = ?", packet.SourceIP).Where("source_port = ?", packet.SourcePort).Where("dest_ip = ?", packet.DestIP).Where("dest_port = ?", packet.DestPort).Order("timestamp asc").First(&firstOccurence)

			tmpFlow := Flow{
				StartTime:  firstOccurence.Timestamp,
				EndTime:    packet.Timestamp,
				SourceIP:   packet.SourceIP,
				SourcePort: packet.SourcePort,
				DestIP:     packet.DestIP,
				DestPort:   packet.DestPort,
			}

			// Append the values to flows
			flows = append(flows, tmpFlow)
		}
	}
	return flows
}
