package handler

import (
	"fmt"

	"github.com/yagikota/network_simulation/model"
	"github.com/yagikota/network_simulation/utils"
)

func ArriveHandler(currentEvent *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig, counter *model.Counter) {
	counter.PacketNum += 1
	// next arrive event
	table.AddEvent(model.ArrivePacket, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))

	if !s.IsBusy() {
		// make the server busy
		s.Use(1)
		table.AddEvent(model.FinishService, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))
		return
	}

	if len(queue.Data) >= queue.Capacity {
		fmt.Println("packet loss")
		counter.PacketLossNum += 1
		return
	}
	queue.Add(model.ArrivePacket, currentEvent.StartTime)
	counter.TotalQueueNum += 1
}
