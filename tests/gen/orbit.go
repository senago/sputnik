package gen

import "github.com/senago/sputnik/internal/domain"

func RandOrbit() domain.Orbit {
	return domain.Orbit{
		ID:       domain.NewOrbitID(),
		Name:     RandString(),
		HeightKm: int64(RandIntInRange(100, 1000)),
	}
}
