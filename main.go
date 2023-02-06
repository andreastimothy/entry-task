package main

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/database"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/server"
)

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Failed to Connect to Database")
	}
	server.Init()
}