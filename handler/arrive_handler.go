package handler

import (
	"fmt"
	"time"

	"github.com/yagikota/network_simulation/model"
)

func ArriveHandler( table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig) {
	// next arrive event
	table.AddEvent(model.ArrivePacket, time.Now())

	if !s.IsBusy() {
		// make the server busy
		s.InUse += 1
		s.Idle -= 1
		table.AddEvent(model.FinishService, time.Now())
		return
	}

	if len(queue.Data) > sConf.K {
		fmt.Println("packet loss")
		// TODO: パケットロスカウント
	}
	queue.Add(model.ArrivePacket)
}
