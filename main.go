package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/zeromq/goczmq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"UST-FireOps/adam/adm"
	"UST-FireOps/adam/parse"
	"UST-FireOps/adam/sql"
)

var (
	dbMutex    sync.Mutex
	zeroMQIP   = "127.0.0.1"
	zeroMQPort = 5555
)

const (
	maxPackets = 128
	snapLen    = 1600
)

func main() {
	collectorAddr := fmt.Sprintf("tcp://%s:%d", zeroMQIP, zeroMQPort)
	pushSock, err := goczmq.NewPush(collectorAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Push socket created")

	defer pushSock.Destroy()

	// Monitor the network interface
	networkInterface := "lo"
	handle, err := pcap.OpenLive(networkInterface, snapLen, true, pcap.BlockForever)

	if err != nil {
		log.Fatal("Unable to listen on interface", err)
	}

	defer handle.Close()

	db, err := gorm.Open(sqlite.Open("flows.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Optimizing the database
	db.Exec("PRAGMA journal_mode = MEMORY;")
	db.Exec("PRAGMA synchronous = OFF;")

	db.AutoMigrate(&sql.PacketDetail{})

	// Run the analyzer periodically
	const interval = 60
	ticker := time.NewTicker(interval * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				adm.Analyze(db, pushSock, &dbMutex)
			}
		}
	}()

	// Write the packets to a database
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := make([]*parse.ParsedPacket, 0, maxPackets)

	for packet := range packetSource.Packets() {
		parsedPacket := parse.ParsePacket(packet)
		if parsedPacket != nil {
			packets = append(packets, parsedPacket)
		}

		if len(packets) >= maxPackets {
			if err := sql.InsertPacketsInBatch(db, &dbMutex, packets); err != nil {
				log.Fatalf("unable to insert data to sqlite: %v", err)
			}
			packets = packets[:0]
		}
	}
}
