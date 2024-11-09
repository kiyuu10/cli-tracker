package task_manager

import (
	"encoding/json"
	"io"
	"os"

	"github.com/kiyuu10/cli-tracker/models"
)

func SaveDB(task *models.TaskList) (err error) {
	db, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(db *os.File) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	dataSave, _ := json.MarshalIndent(task, "", " ")

	// Ghi JSON vào file
	_, err = db.Write(dataSave)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDB() (taskDB models.TaskList, err error) {
	dataDb, err := os.Open("data.json")
	if err != nil {
		return taskDB, err
	}
	defer func(data *os.File) {
		err = data.Close()
		if err != nil {
			return
		}
	}(dataDb)

	data, _ := io.ReadAll(dataDb)
	var tasks models.TaskList
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
