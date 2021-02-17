package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	dataDir := "./dataset/"

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatalln(err)
	}

	resultFile, err := os.Create("final.csv")
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		// ignoring .gitkeep file
		if file.Name() == ".gitkeep" {
			continue
		}

		relativeFilename := dataDir + file.Name()

		reader, osFile := getCSVReader(relativeFilename)
		defer osFile.Close()

		header := getCSVNextRecord(reader)
		respondentIndex := getFieldIndex("Respondent", header)
		countryIndex := getFieldIndex("Country", header)
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := [][]string{{"id", "country"}}
			for {
				record := getCSVNextRecord(reader)
				if record == nil {
					break
				}
				data = append(data, []string{record[respondentIndex], record[countryIndex]})
			}

			w := csv.NewWriter(resultFile)
			w.WriteAll(data)
		}()
	}
	fmt.Println("Concatening csv file concurrently..")
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Took: ", duration)
}
