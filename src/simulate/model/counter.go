package model

// The parameters to calculate outputs.
type Counter struct {
	TotalQueueTime float64 // Total waiting time of packets in queue.
	TotalQueueNum  int     // Total numbers of packets in queue.
	PacketNum      int     // Total numbers of packets.
	PacketLossNum  int     // Numbers of packet loss.
}

func NewCounter() *Counter {
	return new(Counter)
}
