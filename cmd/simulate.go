/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yagikota/network_simulation/src/simulate"
)

// simulateCmd represents the simulate command
func newSimulateCmd() *cobra.Command {
	simulateCmd := &cobra.Command{
		Use:   "simulate",
		Short: "simulate M/M/1/K queue.",
		Long:  `This CLI simulates M/M/1/K queue.🐶`,
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

			simulate.Simulate(l, m, k, st, et)

			return nil
		},
	}

	// TODO: 定数を設定ファイルから読み込む
	simulateCmd.Flags().Float64P("lambda", "l", 0.5, "average arrival rate of a packet")
	simulateCmd.Flags().Float64P("myu", "m", 1.0, "average service rate of the server")
	simulateCmd.Flags().IntP("K", "k", 50, "capacity of service(capacity of queue and server)")
	simulateCmd.Flags().Float64P("start_time", "s", 0.0, "the start time of the simulation")
	simulateCmd.Flags().Float64P("end_time", "e", 3000.0, "the end time of the simulation")

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
