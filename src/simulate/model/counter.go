package model

// The parameters to calculate outputs.
type Counter struct {
	TotalQueueTime  float64 // Total waiting time of packets in queue.
	TotalQueueNum   int     // Total numbers of packets in queue.
	TotalServerTime float64 // Total waiting time of packets in the server.
	LastEventTime   float64 // time of the last event.
	PacketNum       int     // Total numbers of packets.
	PacketLossNum   int     // Numbers of packet loss.
}

func NewCounter(lastEventTime float64) *Counter {
	c := new(Counter)
	c.LastEventTime = lastEventTime
	return c
}
