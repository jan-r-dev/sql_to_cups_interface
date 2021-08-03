package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type deviceStruct struct {
	name       string
	ip         string
	brand      string
	model      string
	ppdNeeded  bool
	ppdAddress string
	options    []string
}

func runSQL() {
	urlDB := getUrlDB()
	query := getQuery()

	// Connect to the database
	conn := connDB(urlDB)
	defer conn.Close(context.Background())

	// Query database
	rows := queryDB(query, conn)
	defer rows.Close()

	// Iterate over returned rows
	structifyQuery(rows)
}

func connDB(urlDB string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), urlDB)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	return conn
}

func queryDB(qs string, conn *pgx.Conn) pgx.Rows {
	rows, err := conn.Query(context.Background(), qs)
	if err != nil {
		log.Fatal("Error querying DB:", err)
	}

	return rows
}

func structifyQuery(rows pgx.Rows) {
	//PLACEHOLDER TODO

	for rows.Next() {
		var name, ip, brand, model, ppdAddress string
		var ppdNeeded bool
		var options []string

		err := rows.Scan(&name, &ip, &brand, &model, &ppdNeeded, &ppdAddress, &options)
		if err != nil {
			log.Fatal("Error reading rows:", err)
		}

		fmt.Println(name, ip, brand, model, ppdAddress, ppdNeeded, options)
	}

	// continue from here - declare struct and return it
}

// Provides valid url for connecting to the database based on env variables
func getUrlDB() string {
	myEnv := importEnv()

	url := myEnv["SQL_PREFIX"] + myEnv["SQL_USER"] + ":" + myEnv["SQL_PASSWORD"] + "@" + myEnv["SQL_IP"] + ":" + myEnv["SQL_PORT"] + "/" + myEnv["SQL_DATABASE"]

	return url
}

func getQuery() string {
	return "select name, ip, brand, model, ppd_needed, ppd_address, options from printers_temp"
	//PLACEHOLDER TODO
}
