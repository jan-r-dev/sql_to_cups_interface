package main

import (
	"log"
	"os"
)

func fileWriteTest(input []string) {

	f, err := os.Create("./output/test.sh")
	if err != nil {
		log.Fatal("Unable to create file:", err)
	}
	defer f.Close()

	initShebang(f)

	for _, c := range input {
		byteC := []byte(c + ";")
		_, err := f.Write(byteC)
		if err != nil {
			log.Fatal("Unable to write to file:", err)
		}
	}
}

func initShebang(f *os.File) {
	_, err := f.Write([]byte("#!/usr/bin/env bash\n\n"))
	if err != nil {
		log.Fatal("Unable to write to file:", err)
	}
}
