package floatrange

import (
	"math/rand"
	"time"
)

func nowRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
