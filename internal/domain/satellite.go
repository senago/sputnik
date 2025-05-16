package domain

import "time"

type SatelliteID UUID

type Satellite struct {
	ID          SatelliteID
	Orbit       Orbit
	Position    Position
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

func (s Satellite) SetPosition(pos Position) Satellite {
	s.Position = pos

	return s
}
