package model

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
	StartTime float64
}

func NewEvent(eventType EventType, time float64) *Event {
	return &Event{
		EventType: eventType,
		StartTime: time,
	}
}

func (t *EventsTable) AddEvent(e EventType, st float64) {
	event := &Event{
		EventType: e,
		StartTime: st,
	}
	t.Events = append(t.Events, event)
}

func (q *EventsTable) IsEmpty() bool {
	return len(q.Events) == 0
}

func (t *EventsTable) Peek() *Event {
	if t.IsEmpty() {
		return nil
	}
	e := t.Events[0]
	t.Events = t.Events[1:]
	return e
}
