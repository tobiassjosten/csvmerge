package main

// usage: go run github.com/tobiassjosten/csvmerge input1.csv input2.csv input3.csv > output.csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	for i, filename := range os.Args[1:] {
		records, err := readFile(filename)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", filename, err)
		}

		if i == 0 {
			err = writer.Write(records[0])
			if err != nil {
				log.Fatalf("Error writing header: %v", err)
			}
		}

		err = writer.WriteAll(records[1:])
		if err != nil {
			log.Fatalf("Error writing data: %v", err)
		}
	}
}

func readFile(filename string) ([][]string, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file %s: %v", filename, err)
	}

	return records, nil
}
