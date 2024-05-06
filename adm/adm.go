package adm

import (
	"fmt"
	"slices"
)

type Dependency struct {
	SrcIP    string
	SrcPort  uint16
	DestIP   string
	DestPort uint16
}

// Check if the timeframe of a flow is encompassed
// within the timeframe of another flow
func isTimeBetween(f1, f2 Flow) bool {
	if (f1.StartTime.After(f2.StartTime) || f1.StartTime.Equal(f2.StartTime)) && (f1.EndTime.Before(f2.EndTime) || f1.EndTime.Equal(f2.EndTime)) {
		return true
	}
	return false
}

func destServ(ip string, port uint16) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

// Find depencencies from flows
func FindDependencies(flows []Flow) []Dependency {
	const threshold float32 = 0.5
	var prevInbounds []Flow
	serviceUsage := make(map[string]int)
	dweight := make(map[Dependency]float32)
	var trackedDependencies []Dependency
	var dependencies []Dependency

	// Iterate through all the flows
	for _, flow := range flows {
		// Increment service usage of flow.DestServ by 1
		serviceUsage[flow.DestServ()] += 1

		// Number of inbound flows within the timeframe of the flow
		// Used in shared mode
		var n float32
		for _, tmpFlow := range flows {
			if tmpFlow != flow {
				if isTimeBetween(flow, tmpFlow) {
					n += 1.0
				}
			}
		}

		// Iterate through previous flow records
		for _, pfl := range prevInbounds {
			if isTimeBetween(flow, pfl) {
				tmpDependency := Dependency{
					SrcIP:    pfl.DestIP,
					SrcPort:  pfl.DestPort,
					DestIP:   flow.DestIP,
					DestPort: flow.DestPort,
				}

				if n > 1 {
					dweight[tmpDependency] += 1.0 / n
				} else {
					dweight[tmpDependency] += 1.0
				}

				// To avoid duplication
				if !slices.Contains(trackedDependencies, tmpDependency) {
					trackedDependencies = append(trackedDependencies, tmpDependency)
				}
			}
		}

		// Append flow to prevInbounds
		prevInbounds = append(prevInbounds, flow)
	}

	for _, dependency := range trackedDependencies {
		aDestServ := destServ(dependency.SrcIP, dependency.SrcPort) // A's Destination Server

		if dweight[dependency]/float32(serviceUsage[aDestServ]) >= threshold {
			dependencies = append(dependencies, dependency)
		}
	}

	return dependencies
}
