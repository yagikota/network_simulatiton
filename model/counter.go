package model

type Counter struct {
	TotalQueueTime float64 // packetのQueueでのtotal待ち時間 TotalQueueTime / simulationTimeでスステム内平均パケット数を求める
	TotalQueueNum  int     // queueに入ったpacket数の総和
	PacketNum      int     // packetの総数
	PacketLossNum  int     // packet loss の回数
}

func NewCounter() *Counter {
	return new(Counter)
}
