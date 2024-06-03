package adm

import (
	"fmt"
	"time"

	"UST-FireOps/adam/sql"

	"gorm.io/gorm"
)

// Flow is defined as a tuple containing
// (StartTime, EndTime, SrcIP, SrcPort, DestinationIP, DestinationPort)
// It represents when a request started and when it ended together with IP and Port details
type Flow struct {
	StartTime time.Time
	EndTime   time.Time
	SrcIP     string
	SrcPort   uint16
	DestIP    string
	DestPort  uint16
}

// Concatenates the IP and Port of the destination service
func (f Flow) DestServ() string {
	return fmt.Sprintf("%s:%d", f.DestIP, f.DestPort)
}

// CreateFlow takes a list of packets (stored in SQLite inmemory database)
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
		db.Model(&sql.PacketDetail{}).Where("src_ip = ?", packet.SrcIP).Where("src_port = ?", packet.SrcPort).Where("dest_ip = ?", packet.DestIP).Where("dest_port = ?", packet.DestPort).Where("timestamp > ?", packet.Timestamp).Count(&count)

		// '0', if there doesn't exists a packet having a greater timestamp than the current one
		if count == 0 {
			// Find the first occurence
			db.Where("src_ip = ?", packet.SrcIP).Where("src_port = ?", packet.SrcPort).Where("dest_ip = ?", packet.DestIP).Where("dest_port = ?", packet.DestPort).Order("timestamp asc").First(&firstOccurence)

			tmpFlow := Flow{
				StartTime: firstOccurence.Timestamp,
				EndTime:   packet.Timestamp,
				SrcIP:     packet.SrcIP,
				SrcPort:   packet.SrcPort,
				DestIP:    packet.DestIP,
				DestPort:  packet.DestPort,
			}

			// Append the values to flows
			flows = append(flows, tmpFlow)
		}
	}
	return flows
}
