package model

import "fmt"

// The input parameters for the simulation.
type SimulationConfig struct {
	Lambda    float64 // Average arrival rate of a packet.
	Myu       float64 // Average service rate of the server
	K         int     // Capacity of service(capacity of queue and server).
	StartTime float64 // The start time of the simulation.
	EndTime   float64 // The end time of the simulation.
}

func NewSimulationConfig(lambda, myu float64, k int, startTime, endTime float64) *SimulationConfig {
	conf := &SimulationConfig{
		Lambda:    lambda,
		Myu:       myu,
		K:         k,
		StartTime: startTime,
	}
	conf.EndTime = conf.StartTime + endTime
	return conf
}

func (s *SimulationConfig) PrintConfInfo() {
	fmt.Println("lambda", s.Lambda)
	fmt.Println("myu", s.Myu)
	fmt.Println("K", s.K)
	fmt.Println("start time", s.StartTime)
	fmt.Println("end time", s.EndTime)
}
