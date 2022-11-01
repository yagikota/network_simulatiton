package utils

import (
	"math"
)

// TODO: search MD1K theoretical values

// MM1KTheoreticalAveragePackets returns the theoretical average number of packets in the system.
func MD1KTheoreticalAveragePackets(lambda, myu float64, k int) float64 {
	float64K := float64(k)
	if lambda == myu {
		return float64K / 2
	}
	rho := lambda / myu
	x := rho / (1 - rho)
	y := (float64K + 1) * math.Pow(rho, float64K+1) / (1 - math.Pow(rho, float64K+1))
	return x - y
}

// MD1KTheoreticalAverageDelay returns the theoretical average delay in the system.
func MD1KTheoreticalAverageDelay(lambda, myu float64, k int) float64 {
	ap := MD1KTheoreticalAveragePackets(lambda, myu, k)
	plr := MD1KTheoreticalPacketLossRate(lambda, myu, k)
	effectiveArrivalRate := lambda * (1 - plr)
	return ap / effectiveArrivalRate
}

// MD1KTheoreticalPacketLossRate returns the theoretical packet loss rate.
func MD1KTheoreticalPacketLossRate(lambda, myu float64, k int) float64 {
	float64K := float64(k)
	if lambda == myu {
		return 1 / (float64K + 1)
	}
	rho := lambda / myu
	return (1 - rho) * math.Pow(rho, float64K) / (1 - math.Pow(rho, float64K+1))
}
