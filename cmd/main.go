package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// simple
// read input
// normalize result
// room for additional processing (later, stage 2)
// write to db
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	records := readCsvFile("../budget_data/2022-01-25_transaction_download.csv")
	fmt.Println(records)
}
