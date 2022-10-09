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

func (q *Queue) Add(t EventType) {
	q.Data = append(q.Data, NewEvent(t))
}

func (q *Queue) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q *Queue) Peek() *Event {
	if q.IsEmpty() {
		return nil
	}
	e := q.Data[0]
	q.Data = q.Data[1:]
	return e
}
