package main

import (
	"fmt"
	"os"
)

func main() {
	// Input of the query as last argument
	filename := os.Args[int(len(os.Args))-1]
	// Retrieve query
	query := getQuery(filename)

	// Retrieve and process devices
	devices := runSQLMain(query)

	// Generate commands
	commands := createCommands(devices)

	fmt.Println(commands)

	fileWriteTest(commands)
}

// go run .\main.go .\sql.go .\helpers.go .\cmds.go .\files.go testquery.sql
