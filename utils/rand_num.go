package utils

import (
	"math"
	"math/rand"
)

func ExpRand(lambda float64) float64 {
	return math.Log(1.0 - rand.Float64()/lambda)
}
