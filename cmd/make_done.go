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

// makeDoneCmd represents the makeDone command
var makeDoneCmd = &cobra.Command{
	Use:   "make-done",
	Short: "make task done",
	Long:  "make task done",
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
				if item.Status == models.TaskStatusDone {
					return
				}
				item.Status = models.TaskStatusDone
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
	rootCmd.AddCommand(makeDoneCmd)
	makeDoneCmd.Flags().Int32("id", 0, "id")
}
