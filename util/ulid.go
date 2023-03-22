package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() ulid.ULID {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	return ulid.MustNew(ulid.MaxTime()-ms, entropy)
}
