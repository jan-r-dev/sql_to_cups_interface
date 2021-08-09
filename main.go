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
		inserts := getInserts(devices)

		fmt.Println(inserts)
	} else if io == "output" {
		filename := "testquery.sql"

		query := getSelect(filename)

		devices := runSQLSelect(query)

		commands := createCommands(devices)

		fmt.Println(commands)

		fileWriteTest(commands)
	} else {
		fmt.Println("Invalid choice. Exiting..")
	}
}

// go run .\main.go .\sql.go .\helpers.go .\cmds.go .\files.go testquery.sql
