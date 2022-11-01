/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/yagikota/network_simulation/src/simulate"
)

// simulateCmd represents the simulate command
func newSimulateCmd() *cobra.Command {
	simulateCmd := &cobra.Command{
		Use:   "simulate",
		Short: "simulate queue.",
		Long:  `This CLI simulates queue.🐶`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("simulate start!")
			l, err := cmd.Flags().GetFloat64("lambda")
			if err != nil {
				return err
			}
			m, err := cmd.Flags().GetFloat64("myu")
			if err != nil {
				return err
			}
			k, err := cmd.Flags().GetInt("K")
			if err != nil {
				return err
			}
			st, err := cmd.Flags().GetFloat64("start_time")
			if err != nil {
				return err
			}
			et, err := cmd.Flags().GetFloat64("end_time")
			if err != nil {
				return err
			}
			qt, err := cmd.Flags().GetInt("queue_type")
			if err != nil {
				return err
			}
			simulate.Simulate(l, m, k, st, et, qt)

			return nil
		},
	}

	// load default input params from env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("could not open env file: %v", err)
	}
	l, _ := strconv.ParseFloat(os.Getenv("DEFAULT_LAMBDA"), 64)
	myu, _ := strconv.ParseFloat(os.Getenv("DEFAULT_MYU"), 64)
	k, _ := strconv.Atoi(os.Getenv("DEFAULT_K"))
	startTime, _ := strconv.ParseFloat(os.Getenv("DEFAULT_START_TIME"), 64)
	endTime, _ := strconv.ParseFloat(os.Getenv("DEFAULT_END_TIME"), 64)
	queueType, _ := strconv.Atoi(os.Getenv("DEFAULT_QUEUE_TYPE"))

	simulateCmd.Flags().Float64P("lambda", "l", l, "average arrival rate of a packet")
	simulateCmd.Flags().Float64P("myu", "m", myu, "average service rate of the server")
	simulateCmd.Flags().IntP("K", "k", k, "capacity of service(capacity of queue and server)")
	simulateCmd.Flags().Float64P("start_time", "s", startTime, "the start time of the simulation")
	simulateCmd.Flags().Float64P("end_time", "e", endTime, "the end time of the simulation")
	simulateCmd.Flags().IntP("queue_type", "q", queueType, "the type of queue(MM1K:0, MD1K: 1)")

	return simulateCmd
}

func init() {
	rootCmd.AddCommand(newSimulateCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
