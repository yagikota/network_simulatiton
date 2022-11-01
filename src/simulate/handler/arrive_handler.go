package handler

import (
	"github.com/yagikota/network_simulation/src/simulate/model"
	"github.com/yagikota/network_simulation/src/simulate/utils"
)

func ArriveHandler(currentEvent *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig, counter *model.Counter) {
	counter.TotalPacketNum += 1
	// next arrive event
	table.AddEvent(model.ArrivePacket, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))

	if !s.IsBusy() {
		// make the server busy
		s.Use(1)
		table.AddEvent(model.FinishService, currentEvent.StartTime+utils.ExpRand(sConf.Myu))
		return
	}

	if len(queue.Data) >= queue.Capacity {
		counter.PacketLossNum += 1
		return
	}
	queue.Add(model.ArrivePacket, currentEvent.StartTime)
}
