package utils

import (
	"time"

	"math/rand"
)

func RandRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
