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

// makeInProgressCmd represents the makeInProgress command
var makeInProgressCmd = &cobra.Command{
	Use:   "make-in-progress",
	Short: "make task in progress",
	Long:  "make task in progress",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			id, _     = cmd.Flags().GetInt32("id")
			taskInfos = make([]models.TaskInfo, 0)
		)

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, item := range tasks.Tasks {
			if item.ID == id {
				if item.Status == models.TaskStatusInProgress {
					return
				}
				item.Status = models.TaskStatusInProgress
			}
			taskInfos = append(taskInfos, item)
		}
		tasks.Tasks = taskInfos

		err = task_manager.SaveDB(&tasks)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(makeInProgressCmd)
	makeInProgressCmd.Flags().Int32("id", 0, "id")
	_ = makeInProgressCmd.MarkFlagRequired("id")
}
