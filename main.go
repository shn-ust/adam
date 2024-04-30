package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"UST-FireOps/adam/parse"
	"UST-FireOps/adam/sql"
)

func main() {
	const snapLen = 262144

	networkInterface := "eth0"
	handle, err := pcap.OpenLive(networkInterface, snapLen, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	// if err := handle.SetBPFFilter("port 5000 or port 5001"); err != nil {
	// 	panic(err)
	// }

	defer handle.Close()

	db, err := gorm.Open(sqlite.Open("flows.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		parsedPacket := parse.ParsePacket(packet)
		if !sql.InsertPacket(parsedPacket, db) {
			panic("Unable to insert data!")
		}
	}
}
