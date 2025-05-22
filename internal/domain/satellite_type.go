package domain

type SatelliteType = string

const (
	SatelliteTypeResourceP = "resource-p"
	SatelliteTypeKanopus   = "kanopus"
	SatelliteTypeKondor    = "kondor"
)

func AllSatelliteType() []SatelliteType {
	return []SatelliteType{
		SatelliteTypeResourceP,
		SatelliteTypeKanopus,
		SatelliteTypeKondor,
	}
}
