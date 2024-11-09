/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/kiyuu10/cli-tracker/models"
	"github.com/kiyuu10/cli-tracker/task_manager"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  "add task",
	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		if description == "" {
			fmt.Println("Please enter description")
			return
		}

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		var task = models.TaskInfo{
			ID:          tasks.CurrentId + 1,
			Description: description,
			Status:      models.TaskStatusTodo,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}
		tasks.Tasks = append(tasks.Tasks, task)
		tasks.CurrentId = tasks.CurrentId + 1

		err = task_manager.SaveDB(&tasks)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("description", "", "description")
	_ = addCmd.MarkFlagRequired("description")
}
