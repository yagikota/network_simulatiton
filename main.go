package main

import (
	"math"
	"math/rand"
	"time"
)

type eventType int

const (
	endSimulation eventType = iota // 0
	arrivePacket                   // 1
	endService                     // 2
)

type inputParam struct {
	Lambda float64
	Myu    float64
	K      int
}

type eventsTable struct {
	Events []*event
}

type event struct {
	eventType eventType
	startTime time.Time
	otherInfo otherInfo
}

type otherInfo struct {
}

func initInputParam() inputParam {
	return inputParam{
		Lambda: 0.2, // 0.2
		Myu:    0.3, // 0.3
		K:      50,  // 50
	}
}

func expRand(lambda float64) float64 {
	return math.Log(1.0 - rand.Float64()/lambda)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 最初のイベントを事象表に登録
	firstEvent := &event{
		eventType: 1,
		startTime: time.Now(),
	}
	eventsTable := new(eventsTable)
	eventsTable.Events = append(eventsTable.Events, firstEvent)

	// 待ち行列を用意
	var queue []*event

	// TODO: シミュレーション開始
}
