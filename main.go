package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	dbUrl := getSqlUrl()

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	defer conn.Close(context.Background())

	// Query database
	rows, err := conn.Query(context.Background(), "select name from printers_temp")
	if err != nil {
		log.Fatal("Error querying DB:", err)
	}
	defer rows.Close()

	// Iterate over returned rows
	for rows.Next() {
		var s string

		err := rows.Scan(&s)
		if err != nil {
			log.Fatal("Error reading rows:", err)
		}

		fmt.Println(s)
	}

}

// Variables needed to connect to the database are stored in a .env file
// This function exports a map with all the needed env variable
func importEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return myEnv
}

// Provides valid url for connecting to the database based on env variables
func getSqlUrl() string {
	myEnv := importEnv()

	url := myEnv["SQL_PREFIX"] + myEnv["SQL_USER"] + ":" + myEnv["SQL_PASSWORD"] + "@" + myEnv["SQL_IP"] + ":" + myEnv["SQL_PORT"] + "/" + myEnv["SQL_DATABASE"]

	return url
}
