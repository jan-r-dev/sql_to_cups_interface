package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/jackc/pgx/v4"
)

type deviceStruct struct {
	name       string
	ip         string
	brand      string
	model      string
	ppdNeeded  bool
	ppdType    string
	ppdAddress string
	options    []string
}

func runSQLMain(query string) []deviceStruct {
	urlDB := getUrlDB()

	// Connect to the database
	conn := connDB(urlDB)
	defer conn.Close(context.Background())

	// Query database
	rows := queryDB(query, conn)
	defer rows.Close()

	// Iterate over returned rows
	devices := structifyQueryResult(rows)

	return devices
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

func structifyQueryResult(rows pgx.Rows) []deviceStruct {
	devices := []deviceStruct{}

	for rows.Next() {
		var name, ip, brand, model, ppdAddress, ppdType string
		var ppdNeeded bool
		var options []string

		err := rows.Scan(&name, &ip, &brand, &model, &ppdNeeded, &ppdType, &ppdAddress, &options)
		if err != nil {
			log.Fatal("Error reading rows:", err)
		}

		ds := deviceStruct{
			name:       name,
			ip:         ip,
			brand:      brand,
			model:      model,
			ppdNeeded:  ppdNeeded,
			ppdType:    ppdType,
			ppdAddress: ppdAddress,
			options:    options,
		}

		devices = append(devices, ds)
	}

	return devices
}

// Provides valid url for connecting to the database based on env variables
func getUrlDB() string {
	myEnv := importEnv()

	url := myEnv["SQL_PREFIX"] + myEnv["SQL_USER"] + ":" + myEnv["SQL_PASSWORD"] + "@" + myEnv["SQL_IP"] + ":" + myEnv["SQL_PORT"] + "/" + myEnv["SQL_DATABASE"]

	return url
}

// Retrieves a query from the given file in the queries folder
func getQuery(filename string) string {
	bOut, err := ioutil.ReadFile("./queries/" + filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	output := string(bOut)

	return output
}
