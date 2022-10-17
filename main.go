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
	rand.Seed(time.Now().UnixNano())

	// ----- BEGIN initialization -----
	simulationConf := model.NewSimulationConfig()
	// register the first event on table.
	eventsTable := new(model.EventsTable)
	firstEvent := model.NewEvent(model.ArrivePacket, 0.0)
	eventsTable.Events = append(eventsTable.Events, firstEvent)
	server := model.NewServer(1)
	queue := model.NewQueue(simulationConf.K - server.Capacity)
	counter := model.NewCounter()
	// ----- END initialization -----

	// ----- BEGIN simulation -----
	var currentEvent *model.Event
	for {
		// pop the event of the nearest future
		if eventsTable.IsEmpty() {
			break
		}
		sort.Slice(eventsTable.Events, func(i, j int) bool {
			return eventsTable.Events[i].StartTime < eventsTable.Events[j].StartTime
		})
		currentEvent = eventsTable.Peek()

		if currentEvent.StartTime > simulationConf.EndTime {
			break
		}

		switch currentEvent.EventType {
		case model.ArrivePacket:
			handler.ArriveHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		case model.FinishService:
			handler.FinishHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		}
	}
	// ----- END simulation -----

	// ----- BEGIN report -----
	fmt.Println("----- Report -----")
	tqt := counter.TotalQueueTime
	l := tqt / currentEvent.StartTime
	fmt.Println("average packets numbers in queue")
	fmt.Println(l)
	w := tqt / float64(counter.TotalQueueNum)
	fmt.Println("average delay of packets in queue")
	fmt.Println(w)
	plr := float64(counter.PacketLossNum) / float64(counter.PacketNum)
	fmt.Println("packets loss rate")
	fmt.Println(plr)
	// ----- END report -----
}
