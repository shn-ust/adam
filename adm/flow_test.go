package adm

import (
	"testing"
	"time"
)

// {TimeStamp:2024-04-23 07:54:20.343191 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}
// {TimeStamp:2024-04-23 07:54:20.34324 +0000 UTC SourceIP:10.0.0.6 SourcePort:8080(http-alt) DestinationIP:10.0.0.4 DestinationPort:58776}
// {TimeStamp:2024-04-23 07:54:20.344677 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}
// {TimeStamp:2024-04-23 07:54:20.344678 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}
// {TimeStamp:2024-04-23 07:54:20.344723 +0000 UTC SourceIP:10.0.0.6 SourcePort:8080(http-alt) DestinationIP:10.0.0.4 DestinationPort:58776}
// {TimeStamp:2024-04-23 07:54:20.346122 +0000 UTC SourceIP:10.0.0.6 SourcePort:49530 DestinationIP:10.0.0.5 DestinationPort:5432(postgresql)}
// {TimeStamp:2024-04-23 07:54:20.347336 +0000 UTC SourceIP:10.0.0.5 SourcePort:5432(postgresql) DestinationIP:10.0.0.6 DestinationPort:49530}
// {TimeStamp:2024-04-23 07:54:20.347346 +0000 UTC SourceIP:10.0.0.6 SourcePort:49530 DestinationIP:10.0.0.5 DestinationPort:5432(postgresql)}
// {TimeStamp:2024-04-23 07:54:20.347507 +0000 UTC SourceIP:10.0.0.6 SourcePort:8080(http-alt) DestinationIP:10.0.0.4 DestinationPort:58776}
// {TimeStamp:2024-04-23 07:54:20.348129 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}
// {TimeStamp:2024-04-23 07:54:20.348129 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}
// {TimeStamp:2024-04-23 07:54:20.348168 +0000 UTC SourceIP:10.0.0.6 SourcePort:8080(http-alt) DestinationIP:10.0.0.4 DestinationPort:58776}
// {TimeStamp:2024-04-23 07:54:20.348569 +0000 UTC SourceIP:10.0.0.4 SourcePort:58776 DestinationIP:10.0.0.6 DestinationPort:8080(http-alt)}

// [Flow(start_time='2024-04-23 07:54:20.343191',
//
//	     end_time='2024-04-23 07:54:20.348569',
//	     sourceIP='10.0.0.4',
//	     sourcePort='58776',
//	     destinationIP='10.0.0.6',
//	     destinationPort='8080'),
//	Flow(start_time='2024-04-23 07:54:20.343240',
//	     end_time='2024-04-23 07:54:20.348168',
//	     sourceIP='10.0.0.6',
//	     sourcePort='8080',
//	     destinationIP='10.0.0.4',
//	     destinationPort='58776}'),
//	Flow(start_time='2024-04-23 07:54:20.346122',
//	     end_time='2024-04-23 07:54:20.347346',
//	     sourceIP='10.0.0.6',
//	     sourcePort='49530',
//	     destinationIP='10.0.0.5',
//	     destinationPort='5432'),
//	Flow(start_time='2024-04-23 07:54:20.347336',
//	     end_time='2024-04-23 07:54:20.347336',
//	     sourceIP='10.0.0.5',
//	     sourcePort='5432',
//	     destinationIP='10.0.0.6',
//	     destinationPort='49530}')]
type Flow struct {
	timeStamp  time.Time
	sourceIP   string
	sourcePort uint16
	destiIP    string
	destPort   uint16
}

func convertToTime(timestamp string) time.Time {
	layout := "2006-01-02 15:04:05.999999"
	converted, err := time.Parse(layout, timestamp)

	if err != nil {
		panic(err)
	}

	return converted
}

func TestCreateFlow(t *testing.T) {
	parsedPackets := `
2024-04-23 07:54:20.343191,10.0.0.4,58776,10.0.0.6,8080
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
2024-04-23 07:54:20.348569,10.0.0.4,58776,10.0.0.6,8080
`
	// flows := []Flow{
	// 	{
	// 		timeStamp: convertToTime("2024-04-23 07:54:20.343191"),
	// 		sourceIP: "10.0.0.4",
	// 		sourcePort: uint16(58776),
	// 	},
	// }
}
