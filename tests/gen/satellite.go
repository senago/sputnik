package gen

import "github.com/senago/sputnik/internal/domain"

func RandSatellite() domain.Satellite {
	return domain.Satellite{
		ID:    domain.NewSatelliteID(),
		Orbit: RandOrbit(),
		Position: domain.NewPosition(
			float32(RandIntInRange(0, 100)),
			float32(RandIntInRange(0, 100)),
		),
		Name:        RandString(),
		Description: RandString(),
		Type:        RandString(),
	}
}
