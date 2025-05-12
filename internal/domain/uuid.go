package domain

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

type UUID = uuid.UUID

func NewUUID(ts time.Time) UUID {
	return uuid.Must(uuid.NewV7AtTime(ts))
}

func UUIDFromString(s string) UUID {
	return uuid.Must(uuid.FromString(s))
}
