package model

import "time"

type SimulationConfig struct {
	Lambda    float64
	Myu       float64
	K         int
	StartTime time.Time
	EndTime   time.Time
}

func NewSimulationConfig() *SimulationConfig {
	conf := &SimulationConfig{
		Lambda:    0.2, // 0.2
		Myu:       0.3, // 0.3
		K:         50,  // 50
		StartTime: time.Now(),
	}
	conf.EndTime = conf.StartTime.Add(10 * time.Second)
	return conf
}
