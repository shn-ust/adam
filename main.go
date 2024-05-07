package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"UST-FireOps/adam/adm"
	"UST-FireOps/adam/parse"
	"UST-FireOps/adam/sql"
)

var dbMutex sync.Mutex

func analyze(db *gorm.DB) {
	fmt.Println("Finding dependency")

	dbMutex.Lock()
	defer dbMutex.Unlock()

	var records []sql.PacketDetail

	if err := db.Find(&records).Error; err != nil {
		log.Fatal("Unable to list records:", err)
	}

	flows := adm.CreateFlow(db)
	dependencies := adm.FindDependencies(flows)

	go func() {
		for _, dependency := range dependencies {
			trueDependency := adm.CheckStatus(dependency)
			if trueDependency {
				fmt.Println(dependency)
			}
		}
	}()

	if err := db.Delete(&records).Error; err != nil {
		log.Fatal("Unable to delete records: ", err)
	}
}

func main() {
	const snapLen = 1600

	networkInterface := "lo"
	handle, err := pcap.OpenLive(networkInterface, snapLen, true, pcap.BlockForever)

	if err != nil {
		log.Fatal("Unable to listen on interface", err)
	}

	// if err := handle.SetBPFFilter("port 5000 or port 5001"); err != nil {
	// 	panic(err)
	// }

	defer handle.Close()

	db, err := gorm.Open(sqlite.Open("flows.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Run the analyzer periodically
	const interval = 15
	ticker := time.NewTicker(interval * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				analyze(db)
			}
		}
	}()

	// Write the packets to a database
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parsedPacket := parse.ParsePacket(packet)

		if !sql.InsertPacket(parsedPacket, db, &dbMutex) {
			log.Fatal("Unable to insert data!")
		}
	}
}
