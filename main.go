package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/yagikota/network_simulation/handler"
	"github.com/yagikota/network_simulation/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// ----- initialization -----
	simulationConf := model.NewSimulationConfig()
	// register the first event on table.
	eventsTable := new(model.EventsTable)
	firstEvent := model.NewEvent(model.EventType(1), 0.0)
	eventsTable.Events = append(eventsTable.Events, firstEvent)
	// test用
	secondEvent := model.NewEvent(model.EventType(2), 1.0)
	eventsTable.Events = append(eventsTable.Events, secondEvent)
	queue := model.NewQueue(simulationConf.K)
	server := model.NewServer()

	// ----- start simulation -----
	for {
		// pop the event of the nearest future
		if eventsTable.IsEmpty() {
			fmt.Println("finish simulation")
			break
		}
		sort.Slice(eventsTable.Events, func(i, j int) bool {
			return eventsTable.Events[i].StartTime < eventsTable.Events[j].StartTime
		})
		currentEvent := eventsTable.Peek()



		switch currentEvent.EventType {
		case model.ArrivePacket:
			fmt.Println("arrive a packet")
			handler.ArriveHandler(currentEvent, eventsTable, queue, server, simulationConf)
		case model.FinishService:
			fmt.Println("end service")
			handler.FinishHandler(currentEvent, eventsTable, queue, server, simulationConf)
		}
		if currentEvent.StartTime > simulationConf.EndTime {
			fmt.Println("finish simulation")
			break
		}
	}
}
