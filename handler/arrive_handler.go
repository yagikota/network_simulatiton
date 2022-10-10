package handler

import (
	"fmt"

	"github.com/yagikota/network_simulation/model"
	"github.com/yagikota/network_simulation/utils"
)

func ArriveHandler(currentEvent *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig) {
	// next arrive event
	table.AddEvent(model.ArrivePacket, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))

	if !s.IsBusy() {
		// make the server busy
		s.InUse += 1
		s.Idle -= 1
		table.AddEvent(model.FinishService, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))
		return
	}

	if len(queue.Data) > sConf.K {
		fmt.Println("packet loss")
		// TODO: パケットロスカウント
		return
	}
	queue.Add(model.ArrivePacket, currentEvent.StartTime)
}
