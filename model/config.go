package model

type SimulationConfig struct {
	Lambda    float64
	Myu       float64
	K         int
	StartTime float64
	EndTime   float64
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
