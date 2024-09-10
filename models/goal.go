package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Goal struct {
	ID          int64
	Name        string
	Description string
	DateTime    time.Time
}

var file_location = "data/goals.txt"

func (e Goal) Save() error {
	// create the file
	f, err := os.Create(file_location)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	// covert the json data to bytes
	goalsBytes, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// write the file
	_, err = f.Write(goalsBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Data written successfully!")
	return err

}
