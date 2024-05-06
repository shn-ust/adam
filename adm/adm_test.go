package adm

import (
	"slices"
	"testing"
)

func TestFindDependencies(t *testing.T) {
	flows := []Flow{
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

	got := FindDependencies(flows)
	want := Dependency{
		SrcIP:    "10.0.0.6",
		SrcPort:  uint16(8080),
		DestIP:   "10.0.0.5",
		DestPort: uint16(5432),
	}

	if !slices.Contains(got, want) {
		t.Errorf("Dependency not found in array! got %+v, want %+v", got, want)
	}
}
