package handler

import (
	"github.com/yagikota/network_simulation/src/simulate/model"
	"github.com/yagikota/network_simulation/src/simulate/utils"
)

func FinishHandler(currentEvent *model.Event, table *model.EventsTable, queue *model.Queue, s *model.Server, sConf *model.SimulationConfig, counter *model.Counter) {
	if queue.IsEmpty() {
		s.Free(1)
		return
	}

	// サーバーの状態はそのままにして後続の処理をする
	event := queue.Peek()
	if event != nil {
		// calculate total queue time
		counter.TotalQueueTime += currentEvent.StartTime - event.StartTime
	}
	// fmt.Println("counter.TotalQueueTime",counter.TotalQueueTime)
	table.AddEvent(model.FinishService, currentEvent.StartTime+utils.ExpRand(sConf.Myu))
}
