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

// Add adds event into queue.
func (t *EventsTable) AddEvent(e EventType, time float64) {
	event := &Event{
		EventType: e,
		StartTime: time,
	}
	t.Events = append(t.Events, event)
}

// IsEmpty returns true if the event table is empty.
func (q *EventsTable) IsEmpty() bool {
	return len(q.Events) == 0
}

// Peek removes an event from the left side of the event table and returns the event.
func (t *EventsTable) Peek() *Event {
	if t.IsEmpty() {
		return nil
	}
	e := t.Events[0]
	t.Events = t.Events[1:]
	return e
}
