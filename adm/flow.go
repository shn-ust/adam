package adm

import (
	"time"

	"gorm.io/gorm"
)

type Flow struct {
	StartTime  time.Time
	EndTime    time.Time
	SourceIP   string
	SourcePort uint16
	DestIP     string
	DestPort   uint16
}

func CreateFlow(db *gorm.DB) []Flow {
	return []Flow{}
}
