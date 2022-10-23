package simulate

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/yagikota/network_simulation/src/simulate/handler"
	"github.com/yagikota/network_simulation/src/simulate/model"
)

func Simulate(lambda, myu float64, k int, startTime, endTime float64) {
	rand.Seed(time.Now().UnixNano())
	// ----- BEGIN initialization -----
	simulationConf := model.NewSimulationConfig(lambda, myu, k, startTime, endTime)
	// register the first event on table.
	eventsTable := new(model.EventsTable)
	firstEvent := model.NewEvent(model.ArrivePacket, startTime)
	eventsTable.Events = append(eventsTable.Events, firstEvent)
	server := model.NewServer(1) // In this time, set 1.
	queue := model.NewQueue(simulationConf.K - server.Capacity)
	counter := model.NewCounter(firstEvent.StartTime)
	// ----- END initialization -----

	// ----- BEGIN simulation -----
	var currentEvent *model.Event
	for {
		// Pop the event of the nearest future from event table.
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

		// counter handling
		timeSinceLastEvent := currentEvent.StartTime - counter.LastEventTime
		counter.LastEventTime = currentEvent.StartTime
		counter.TotalQueueTime += float64(len(queue.Data)) * timeSinceLastEvent
		counter.TotalServerTime += float64(server.InUse) * timeSinceLastEvent

		switch currentEvent.EventType {
		case model.ArrivePacket:
			handler.ArriveHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		case model.FinishService:
			handler.FinishHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		}
	}
	// ----- END simulation -----

	// ----- BEGIN report -----
	fmt.Println("----- Input Params -----")
	simulationConf.PrintConfInfo()
	fmt.Println("----- Report -----")
	totalTimeInService := counter.TotalQueueTime + counter.TotalServerTime
	simulateTime := currentEvent.StartTime - simulationConf.StartTime
	l := totalTimeInService / simulateTime
	fmt.Println("average packets numbers in queue:", l)

	w := totalTimeInService / float64(counter.TotalQueueNum)
	fmt.Println("average delay of packets in queue:", w)

	plr := float64(counter.PacketLossNum) / float64(counter.PacketNum)
	// fmt.Println("packets", counter.PacketNum)
	fmt.Println("packets loss rate:", plr)
	// ----- END report -----
}
