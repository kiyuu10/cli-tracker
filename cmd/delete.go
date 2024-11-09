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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	Long:  `delete task that done or no effective`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt32("id")
		if id == 0 {
			fmt.Println("Please enter id")
			return
		}

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		var taskInfos = make([]models.TaskInfo, 0)
		for _, item := range tasks.Tasks {
			if item.ID != id {
				taskInfos = append(taskInfos, item)
			}
		}
		if tasks.CurrentId == id {
			tasks.CurrentId = taskInfos[len(taskInfos)-1].ID
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
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int32("id", 0, "id")
	_ = deleteCmd.MarkFlagRequired("id")
}
