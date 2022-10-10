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
	queue := model.NewQueue(simulationConf.K)
	server := model.NewServer(1)
	counter := model.NewCounter()
	// ----- END initialization -----

	// ----- BEGIN simulation -----
	var currentEvent *model.Event
	for {
		// pop the event of the nearest future
		if eventsTable.IsEmpty() {
			fmt.Println("finish simulation")
			break
		}
		sort.Slice(eventsTable.Events, func(i, j int) bool {
			return eventsTable.Events[i].StartTime < eventsTable.Events[j].StartTime
		})
		currentEvent = eventsTable.Peek()
		if currentEvent.StartTime > simulationConf.EndTime {
			fmt.Println("finish simulation")
			break
		}

		switch currentEvent.EventType {
		case model.ArrivePacket:
			fmt.Println("arrive a packet")
			handler.ArriveHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		case model.FinishService:
			fmt.Println("end service")
			handler.FinishHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		}
	}
	// ----- END simulation -----

	// ----- BEGIN report -----
	tqt := counter.TotalQueueTime
	l := tqt / currentEvent.StartTime
	fmt.Println("average packets numbers in queue")
	fmt.Println(l)
	w := tqt / float64(counter.TotalQueueNum)
	fmt.Println("average delay of packets in queue")
	fmt.Println(w)
	plr := counter.PacketLossNum / counter.PacketNum
	fmt.Println("packets loss rate")
	fmt.Println(plr)
	// ----- END report -----
}
