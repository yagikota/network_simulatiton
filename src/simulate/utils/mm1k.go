package utils

import (
	"math"
	"math/rand"
)

// https://blog.monochromegane.com/blog/2019/10/11/random_number_gen_using_go/
func ExpRand(x float64) float64 {
	r := rand.ExpFloat64() / x
	return r
}

// MM1KTheoreticalAveragePackets returns the theoretical average number of packets in the system.
// https://sites.pitt.edu/~dtipper/2130/2130_Slides4.pdf
func MM1KTheoreticalAveragePackets(lambda, myu float64, k int) float64 {
	float64K := float64(k)
	if lambda == myu {
		return float64K / 2
	}
	rho := lambda / myu
	x := rho / (1 - rho)
	y := (float64K + 1) * math.Pow(rho, float64K+1) / (1 - math.Pow(rho, float64K+1))
	return x - y
}

// MM1KTheoreticalAverageDelay returns the theoretical average delay in the system.
func MM1KTheoreticalAverageDelay(lambda, myu float64, k int) float64 {
	ap := MM1KTheoreticalAveragePackets(lambda, myu, k)
	plr := MM1KTheoreticalPacketLossRate(lambda, myu, k)
	effectiveArrivalRate := lambda * (1 - plr)
	return ap / effectiveArrivalRate
}

// MM1KTheoreticalPacketLossRate returns the theoretical packet loss rate.
func MM1KTheoreticalPacketLossRate(lambda, myu float64, k int) float64 {
	float64K := float64(k)
	if lambda == myu {
		return 1 / (float64K + 1)
	}
	rho := lambda / myu
	return (1 - rho) * math.Pow(rho, float64K) / (1 - math.Pow(rho, float64K+1))
}
