package domain

import "time"

type OrbitID UUID

type Orbit struct {
	ID       OrbitID
	Name     string
	HeightKm int64
}

func NewOrbitID() OrbitID {
	return OrbitID(NewUUID(time.Now()))
}

func (oid OrbitID) String() string {
	return UUID(oid).String()
}
