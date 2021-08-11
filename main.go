package main

import (
	"fmt"
)

func main() {
	fmt.Println("Choose input or output: ")
	var io string
	fmt.Scanln(&io)

	if io == "input" {
		filename := "input.csv"
		file := openFile(filename)
		defer file.Close()

		devices := createDevicesFromFile(file)

		runSQLInserts(devices)

	} else if io == "output" {
		filename := "testquery.sql"

		query := getSelect(filename)

		devices := runSQLSelect(query)
		fmt.Println(devices)

		//commands := createCommands(devices)

		//fileWriteTest(commands)
	} else {
		fmt.Println("Invalid choice. Exiting..")
	}
}
