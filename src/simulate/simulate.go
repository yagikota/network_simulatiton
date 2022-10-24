package simulate

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/yagikota/network_simulation/src/simulate/handler"
	"github.com/yagikota/network_simulation/src/simulate/model"
)

const (
	csvFilePathFromRoot = "./csv_report/report.csv"
)

func Simulate(lambda, myu float64, k int, startTime, endTime float64) {
	rand.Seed(time.Now().UnixNano())
	// ----- BEGIN initialization -----
	simulationConf := model.NewSimulationConfig(lambda, myu, k, startTime, endTime)
	// register the first event on table.
	eventsTable := new(model.EventsTable)
	firstEvent := model.NewEvent(model.ArrivePacket, startTime)
	eventsTable.Events = append(eventsTable.Events, firstEvent)
	server := model.NewServer(1) // In this time, set 1.
	queue := model.NewQueue(simulationConf.K - server.Capacity)
	counter := model.NewCounter(firstEvent.StartTime)
	// ----- END initialization -----

	// ----- BEGIN simulation -----
	var currentEvent *model.Event
	for {
		// Pop the event of the nearest future from event table.
		if eventsTable.IsEmpty() {
			break
		}
		sort.Slice(eventsTable.Events, func(i, j int) bool {
			return eventsTable.Events[i].StartTime < eventsTable.Events[j].StartTime
		})
		currentEvent = eventsTable.Peek()
		if currentEvent.StartTime > simulationConf.EndTime {
			break
		}

		// counter handling
		timeSinceLastEvent := currentEvent.StartTime - counter.LastEventTime
		counter.LastEventTime = currentEvent.StartTime
		counter.TotalQueueTime += float64(len(queue.Data)) * timeSinceLastEvent
		counter.TotalServerTime += float64(server.InUse) * timeSinceLastEvent

		switch currentEvent.EventType {
		case model.ArrivePacket:
			handler.ArriveHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		case model.FinishService:
			handler.FinishHandler(currentEvent, eventsTable, queue, server, simulationConf, counter)
		}
	}
	// ----- END simulation -----

	totalTimeInService := counter.TotalQueueTime + counter.TotalServerTime
	simulateTime := currentEvent.StartTime - simulationConf.StartTime
	averagePackets := totalTimeInService / simulateTime                           // average number of packets in the system
	averageDelay := totalTimeInService / float64(counter.TotalQueueNum)           // average delay in the system
	packetLossRate := float64(counter.PacketLossNum) / float64(counter.PacketNum) // packet loos rate

	// ----- BEGIN report -----
	printReport(counter, simulationConf, averagePackets, averageDelay, packetLossRate)
	// ----- END report -----

	// ----- BEGIN export the report to csv -----
	if err := printCSV(counter, simulationConf, averagePackets, averageDelay, packetLossRate); err != nil {
		log.Fatal(err)
	}
	// ----- END export the report to csv -----
}

func printReport(counter *model.Counter, conf *model.SimulationConfig, averagePackets, averageDelay, packetLossRate float64) {
	fmt.Println(strings.Repeat("-", 5), "Input Params", strings.Repeat("-", 5))
	conf.PrintConfInfo()
	fmt.Println(strings.Repeat("-", 5), "Report", strings.Repeat("-", 5))
	fmt.Println("average packets numbers in the system:", averagePackets)
	fmt.Println("average delay of packets in the system:", averageDelay)
	fmt.Println("packets loss rate:", packetLossRate)
}

// https://zenn.dev/hsaki/books/golang-io-package/viewer/file
func printCSV(counter *model.Counter, conf *model.SimulationConfig, averagePackets, averageDelay, packetLossRate float64) error {
	var f *os.File
	if _, err := os.Stat(csvFilePathFromRoot); err == nil {
		f, err = os.OpenFile(csvFilePathFromRoot, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		log.Println("successfully opened file:", f.Name())

	} else if errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(csvFilePathFromRoot)
		if err != nil {
			return err
		}
		// header
		if _, err := f.WriteString(strings.Join([]string{"lambda", "myu", "K", "L", "W", "Q\n"}, ",")); err != nil {
			return err
		}
		log.Println("successfully created file:", f.Name())
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	w := csv.NewWriter(f)
	defer w.Flush()

	lambda := fmt.Sprintf("%f", conf.Lambda)
	myu := fmt.Sprintf("%f", conf.Myu)
	k := fmt.Sprintf("%d", conf.K)
	ap := strconv.FormatFloat(averagePackets, 'f', -1, 64)
	ad := strconv.FormatFloat(averageDelay, 'f', -1, 64)
	plr := strconv.FormatFloat(packetLossRate, 'f', -1, 64)
	records := []string{lambda, myu, k, ap, ad, plr}
	if err := w.Write(records); err != nil {
		return err
	}

	return nil
}
