package model

// The input parameters for the simulation.
type SimulationConfig struct {
	Lambda    float64 // Average arrival rate of a packet.
	Myu       float64 // Average service rate of the server
	K         int     // Capacity of service(capacity of queue and server).
	StartTime float64 // The start time of the simulation.
	EndTime   float64 // The end time of the simulation.
}

func NewSimulationConfig() *SimulationConfig {
	conf := &SimulationConfig{
		Lambda:    0.2, // 0.2
		Myu:       0.3, // 0.3
		K:         50,  // 50
		StartTime: 0.0,
	}
	conf.EndTime = conf.StartTime + 1000.0
	return conf
}
