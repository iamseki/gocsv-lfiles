package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func getCSVReader(filename string) (*csv.Reader, *os.File) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	return r, csvfile
}

func getCSVNextRecord(r *csv.Reader) []string {
	record, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return record
}

func getFieldIndex(field string, record []string) int {
	for i, v := range record {
		if v == field {
			return i
		}
	}
	return -1
}
