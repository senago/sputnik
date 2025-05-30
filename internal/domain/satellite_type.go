package domain

type SatelliteType = string

const (
	SatelliteTypeResourceP = "РЕСУРС-П"
	SatelliteTypeKanopus   = "КАНОПУС"
	SatelliteTypeKondor    = "КОНДОР"
	SatelliteTypeMeteorM   = "МЕТЕОР-М"
	SatelliteTypeElectroL  = "ЭЛЕКТРО-Л"
)

func AllSatelliteType() []SatelliteType {
	return []SatelliteType{
		SatelliteTypeResourceP,
		SatelliteTypeKanopus,
		SatelliteTypeKondor,
		SatelliteTypeMeteorM,
		SatelliteTypeElectroL,
	}
}
