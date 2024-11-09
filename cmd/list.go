/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kiyuu10/cli-tracker/models"
	"github.com/kiyuu10/cli-tracker/task_manager"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list task",
	Long:  "list task",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			statusArg   string
			statusQuery models.TaskStatus
		)
		if statusArg == "todo" {
			statusQuery = 1
		} else if statusArg == "in-progress" {
			statusQuery = 2
		} else if statusArg == "done" {
			statusQuery = 3
		}

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("======================Tasks:=========================")
		for _, task := range tasks.Tasks {
			if statusQuery == 0 {
				var statusRes string
				switch task.Status {
				case models.TaskStatusTodo:
					statusRes = "todo"
				case models.TaskStatusInProgress:
					statusRes = "in-progress"
				case models.TaskStatusDone:
					statusRes = "done"
				}
				fmt.Printf("\n+ Task %v: %s - status: %s", task.ID, task.Description, statusRes)
			} else if task.Status == statusQuery {
				fmt.Printf("\n+ Task %v: %s - status: %s", task.ID, task.Description, statusArg)
			}
		}
		fmt.Println("\n=====================================================")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
