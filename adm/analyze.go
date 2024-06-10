package adm

import (
	"fmt"
	"log"
	"sync"

	"github.com/zeromq/goczmq"
	"gorm.io/gorm"

	"UST-FireOps/adam/sql"
)

// idFromRecords is used to return the ID of the records
// Used for batch delete
func idFromRecords(records []sql.PacketDetail) []uint {
	var rec []uint
	for _, record := range records {
		rec = append(rec, record.ID)
	}
	return rec
}

// Analyze is used to find the dependencies from the stored packet data
// If dependencies are found, it sends them to the collector using zeromq's push-pull pattern
func Analyze(db *gorm.DB, pushSock *goczmq.Sock, dbMutex *sync.Mutex) {
	log.Println("Finding dependency")

	dbMutex.Lock()
	defer dbMutex.Unlock()

	var records []sql.PacketDetail

	if err := db.Find(&records).Error; err != nil {
		log.Fatal("Unable to list records:", err)
	}

	flows := CreateFlow(db)
	dependencies := FindDependencies(flows)

	for _, dependency := range dependencies {
		trueDependency := CheckStatus(dependency)

		if trueDependency {
			log.Println(dependency)
			dependencyStr := fmt.Sprintf("%s:%d,%s:%d", dependency.SrcIP, dependency.SrcPort, dependency.DestIP, dependency.DestPort)
			err := pushSock.SendFrame([]byte(dependencyStr), 0)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if len(records) > 0 {
		if err := db.Unscoped().Delete(&sql.PacketDetail{}, idFromRecords(records)).Error; err != nil {
			log.Fatal("Unable to delete records: ", err)
		}
	}
}
