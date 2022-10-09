package model

import "time"

type EventType int

const (
	EndSimulation EventType = iota // 0
	ArrivePacket                   // 1
	FinishService                  // 2
)

type EventsSlice []*Event

type EventsTable struct {
	Events EventsSlice
}

type Event struct {
	EventType EventType
	StartTime time.Time
}

func NewEvent(t EventType) *Event {
	return &Event{
		EventType: t,
		StartTime: time.Now(),
	}
}

func (t *EventsTable) AddEvent(e EventType, st time.Time) {
	event := &Event{
		EventType: e,
		StartTime: st,
	}
	t.Events = append(t.Events, event)
}
