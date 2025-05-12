package domain

import "time"

type SatelliteID UUID

type Satellite struct {
	ID          SatelliteID
	Orbit       Orbit
	Name        string
	Description string
	Type        string
}

func NewSatelliteID() SatelliteID {
	return SatelliteID(NewUUID(time.Now()))
}

func SatelliteIDFromString(s string) SatelliteID {
	return SatelliteID(UUIDFromString(s))
}

func (sid SatelliteID) String() string {
	return UUID(sid).String()
}
