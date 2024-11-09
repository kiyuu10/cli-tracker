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
		if len(args) == 0 && args[0] == "" {
			fmt.Println("please enter description of task")
		}

		tasks, err := task_manager.GetTaskDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		var task = models.TaskInfo{
			ID:          tasks.CurrentId + 1,
			Description: args[0],
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
