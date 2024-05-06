package adm

import (
	"UST-FireOps/adam/sql"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func convertToTime(timestamp string) time.Time {
	layout := "2006-01-02 15:04:05.999999"
	converted, err := time.Parse(layout, timestamp)

	if err != nil {
		panic(err)
	}

	return converted
}

func convertToUint16(number string) uint16 {
	converted, err := strconv.ParseUint(number, 10, 16)

	if err != nil {
		panic(err)
	}

	return uint16(converted)

}

// Helper function to populate DB with values for testing
func populateDB(db *gorm.DB) {
	packetsDump := `2024-04-23 07:54:20.343191,10.0.0.4,58776,10.0.0.6,8080
2024-04-23 07:54:20.34324,10.0.0.6,8080,10.0.0.4,58776
2024-04-23 07:54:20.344677,10.0.0.4,58776,10.0.0.6,8080
2024-04-23 07:54:20.344678,10.0.0.4,58776,10.0.0.6,8080
2024-04-23 07:54:20.344723,10.0.0.6,8080,10.0.0.4,58776
2024-04-23 07:54:20.346122,10.0.0.6,49530,10.0.0.5,5432
2024-04-23 07:54:20.347336,10.0.0.5,5432,10.0.0.6,49530
2024-04-23 07:54:20.347346,10.0.0.6,49530,10.0.0.5,5432
2024-04-23 07:54:20.347507,10.0.0.6,8080,10.0.0.4,58776
2024-04-23 07:54:20.348129,10.0.0.4,58776,10.0.0.6,8080
2024-04-23 07:54:20.348129,10.0.0.4,58776,10.0.0.6,8080
2024-04-23 07:54:20.348168,10.0.0.6,8080,10.0.0.4,58776
2024-04-23 07:54:20.348569,10.0.0.4,58776,10.0.0.6,8080`

	splitted := strings.Split(packetsDump, "\n")
	var packetDetails []sql.PacketDetail

	for _, packet := range splitted {
		detail := strings.Split(packet, ",")
		tmpFLow := sql.PacketDetail{
			Timestamp: convertToTime(detail[0]),
			SrcIP:     detail[1],
			SrcPort:   convertToUint16(detail[2]),
			DestIP:    detail[3],
			DestPort:  convertToUint16(detail[4]),
		}
		packetDetails = append(packetDetails, tmpFLow)
	}

	db.AutoMigrate(&sql.PacketDetail{})
	for _, packetDetail := range packetDetails {
		db.Create(&packetDetail)
	}
}

func TestCreateFlow(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	populateDB(db)

	got := CreateFlow(db)
	want := []Flow{
		{
			StartTime: convertToTime("2024-04-23 07:54:20.343191"),
			EndTime:   convertToTime("2024-04-23 07:54:20.348569"),
			SrcIP:     "10.0.0.4",
			SrcPort:   uint16(58776),
			DestIP:    "10.0.0.6",
			DestPort:  uint16(8080),
		},
		{
			StartTime: convertToTime("2024-04-23 07:54:20.343240"),
			EndTime:   convertToTime("2024-04-23 07:54:20.348168"),
			SrcIP:     "10.0.0.6",
			SrcPort:   uint16(8080),
			DestIP:    "10.0.0.4",
			DestPort:  uint16(58776),
		},
		{
			StartTime: convertToTime("2024-04-23 07:54:20.346122"),
			EndTime:   convertToTime("2024-04-23 07:54:20.347346"),
			SrcIP:     "10.0.0.6",
			SrcPort:   uint16(49530),
			DestIP:    "10.0.0.5",
			DestPort:  uint16(5432),
		},
		{
			StartTime: convertToTime("2024-04-23 07:54:20.347336"),
			EndTime:   convertToTime("2024-04-23 07:54:20.347336"),
			SrcIP:     "10.0.0.5",
			SrcPort:   uint16(5432),
			DestIP:    "10.0.0.6",
			DestPort:  uint16(49530),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
