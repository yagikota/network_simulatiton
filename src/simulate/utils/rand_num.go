package utils

import (
	"math/rand"
)

// https://blog.monochromegane.com/blog/2019/10/11/random_number_gen_using_go/
func ExpRand(x float64) float64 {
	r := rand.ExpFloat64() / x
	// r := -math.Log(float64(1.0 - rand.Float64())) / x
	return r
}
