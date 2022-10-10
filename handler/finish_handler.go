package handler

import (
	"github.com/yagikota/network_simulation/model"
	"github.com/yagikota/network_simulation/utils"
)

func FinishHandler(currentTime *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig) {
	if queue.IsEmpty() {
		// make server idle
		s.InUse -= 1
		s.Idle += 1
		return
	}

	// サーバーの状態はそのままにして後続の処理をする
	queue.Peek()
	table.AddEvent(model.FinishService, currentTime.StartTime+utils.ExpRand(sConf.Lambda))
}
