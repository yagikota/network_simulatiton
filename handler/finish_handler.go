package handler

import (
	"time"

	"github.com/yagikota/network_simulation/model"
)

func FinishHandler(table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig) {
	if queue.IsEmpty() {
		s.InUse -= 1
		s.Idle += 1
		return
	}

	// サーバーの状態はそのままにして後続の処理をする
	queue.Peek()
	table.AddEvent(model.EventType(2), time.Now())
}
