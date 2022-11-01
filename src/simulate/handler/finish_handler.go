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

	// queueに入っているeventを取り出してサーバーで処理をする
	_ = queue.Peek()

	var serverTime float64
	switch qt := sConf.QueueType; qt {
	case model.MM1K:
		serverTime = utils.ExpRand(sConf.Myu)
	case model.MD1K:
		serverTime = 1 / sConf.Myu
	default:
		serverTime = utils.ExpRand(sConf.Myu)
	}

	table.AddEvent(model.FinishService, currentEvent.StartTime+serverTime)
}
