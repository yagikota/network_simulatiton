package model

// The parameters to calculate outputs.
type Counter struct {
	TotalQueueTime  float64 // Total waiting time of packets in queue.
	TotalServerTime float64 // Total waiting time of packets in the server.
	LastEventTime   float64 // Time of the last event.
	TotalPacketNum  int     // Total number of packets.
	PacketLossNum   int     // The number of packet loss.
}

func NewCounter(lastEventTime float64) *Counter {
	c := new(Counter)
	c.LastEventTime = lastEventTime
	return c
}
