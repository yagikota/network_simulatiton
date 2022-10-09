package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/yagikota/network_simulation/handler"
	"github.com/yagikota/network_simulation/model"
)

func main() {
	// ----- initialization -----
	rand.Seed(time.Now().UnixNano())
	simulationConf := model.NewSimulationConfig()
	// 最初のイベントを事象表に登録
	firstEvent := model.NewEvent(model.EventType(1))
	eventsTable := new(model.EventsTable)
	eventsTable.Events = append(eventsTable.Events, firstEvent)
	// test用
	secondEvent := model.NewEvent(model.EventType(2))
	eventsTable.Events = append(eventsTable.Events, secondEvent)
	// 待ち行列を用意
	queue := model.NewQueue(simulationConf.K)

	server := model.NewServer()
	// ----- start simulation -----
	for {
		if len(eventsTable.Events) == 0 {
			fmt.Println("finish simulation")
			break
		}
		sort.Slice(eventsTable.Events, func(i, j int) bool {
			return eventsTable.Events[i].StartTime.After(eventsTable.Events[j].StartTime)
		})
		// eventsTable.Events
		currEvent := eventsTable.Events[0]
		eventsTable.Events = eventsTable.Events[1:]

		currentTime := currEvent.StartTime

		switch currEvent.EventType {
		case model.ArrivePacket:
			fmt.Println("arrive a packet")
			handler.ArriveHandler(eventsTable, queue, server, simulationConf)
		case model.FinishService:
			fmt.Println("end service")
			handler.FinishHandler(eventsTable, queue, server, simulationConf)
		}
		if currentTime.After(simulationConf.EndTime) {
			// 終了処理
			fmt.Println("finish simulation")
			break
		}
	}
}
