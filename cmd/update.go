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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task",
	Long:  "update task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt32("id")
		description, _ := cmd.Flags().GetString("description")

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		var taskInfos = make([]models.TaskInfo, 0)
		for _, item := range tasks.Tasks {
			if item.ID == id {
				item.Description = description
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
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Int32("id", 0, "id")
	_ = updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().String("description", "", "description")
	_ = updateCmd.MarkFlagRequired("description")
}
