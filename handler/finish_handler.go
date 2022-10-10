package handler

import (
	"github.com/yagikota/network_simulation/model"
	"github.com/yagikota/network_simulation/utils"
)

func FinishHandler(currentEvent *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig, counter *model.Counter) {
	if queue.IsEmpty() {
		// make server idle
		s.InUse -= 1
		s.Idle += 1
		return
	}

	// サーバーの状態はそのままにして後続の処理をする
	event := queue.Peek()
	// calculate total queue time
	counter.TotalQueueTime += currentEvent.StartTime - event.StartTime
	table.AddEvent(model.FinishService, currentEvent.StartTime+utils.ExpRand(sConf.Lambda))
}
