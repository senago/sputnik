package gen

import (
	"math/rand"

	"github.com/gofrs/uuid/v5"
)

func RandString() string {
	return uuid.Must(uuid.NewV7()).String()
}

func RandIntInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
