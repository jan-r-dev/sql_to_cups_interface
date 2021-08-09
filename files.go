package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
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

func openFile(filename string) *os.File {
	f, err := os.Open("./input/" + filename)
	if err != nil {
		log.Fatal("Unable to open file:", err)
	}

	return f
}

func createDevicesFromFile(f *os.File) []deviceStruct {
	devices := []deviceStruct{}

	fl, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal("Unable to open file:", err)
	}

	for _, line := range fl {
		ppdNeeded := false
		ppdType := ""
		ppdAddress := ""
		options := []string{"placeholder/query/for/options"}

		brand := checkBrand(line[2])

		if brand != "hp" {
			ppdNeeded = true
			ppdAddress = "placeholder/query/for/address"

			if brand == "zebra" {
				ppdType = "m"
			} else {
				ppdType = "P"
			}
		} else if brand == "unknown" {
			ppdType = "null"
			ppdAddress = "null"
			ppdType = "null"
		} else {
			ppdAddress = "null"
			ppdType = "null"
		}

		ds := deviceStruct{
			name:       line[0],
			ip:         line[1],
			brand:      brand,
			model:      line[2],
			ppdNeeded:  ppdNeeded,
			ppdType:    ppdType,
			ppdAddress: ppdAddress,
			options:    options,
		}

		devices = append(devices, ds)
	}

	return devices
}

func checkBrand(m string) string {
	if strings.Contains(strings.ToLower(m), "intermec") {
		return "intermec"
	} else if strings.Contains(strings.ToLower(m), "honeywell") {
		return "honeywell"
	} else if strings.Contains(strings.ToLower(m), "hp ") {
		return "hp"
	} else if strings.Contains(strings.ToLower(m), "zebra") {
		return "zebra"
	} else {
		return "unknown"
	}
}
