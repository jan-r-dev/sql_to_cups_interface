package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	myEnv := importEnv()

	fmt.Println(myEnv)
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
