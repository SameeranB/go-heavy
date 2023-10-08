/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/SameeranB/go-heavy/internal/utils"
	"github.com/spf13/cobra"
)

var duration int
var concurrency int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [path to workflow]",
	Args:  cobra.ExactArgs(1),
	Short: "Run a workflow",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pathToWorkflow := args[0]

		workflowConfig, err := utils.GetWorkflowConfigFromTest(pathToWorkflow)
		if err != nil {
			panic(err)
		}

		err = utils.RunWorkflow(workflowConfig, duration, concurrency)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().IntVarP(&duration, "duration", "d", 10, "Duration of the test in seconds")
	runCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 2, "Number of concurrent workflows to run")
}
