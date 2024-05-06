package adm

type Dependency struct {
	SrcIP   string
	SrcPort uint16
	DstIP   string
	DstPort uint16
}

func FindDependencies(flows []Flow) []Dependency {
	return []Dependency{}
}
