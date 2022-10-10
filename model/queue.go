package model

// https://fodor.org/blog/go-queue-and-stack/
type Queue struct {
	Data []*Event
}

func NewQueue(cap int) *Queue {
	return &Queue{
		Data: make([]*Event, 0, cap),
	}
}

// Add adds event into queue.
func (q *Queue) Add(eventType EventType, time float64) {
	q.Data = append(q.Data, NewEvent(eventType, time))
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.Data) == 0
}

// Peek removes an event from the left side of the queue and returns the event.
func (q *Queue) Peek() *Event {
	if q.IsEmpty() {
		return nil
	}
	e := q.Data[0]
	q.Data = q.Data[1:]
	return e
}
