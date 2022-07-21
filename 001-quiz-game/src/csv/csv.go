package csv

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

// ReadCSVFile loads a CSV file
func ReadCSVFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read input file %v with error %v", filename, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse csv file %v with error %v", filename, err)
	}
	return records
}

func ShuffleData(data [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}
