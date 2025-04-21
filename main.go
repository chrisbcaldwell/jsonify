package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Parse flags
	filePath := flag.String("path", "", "Path of CSV file to convert to JSON")
	flag.Parse()

	// If no file specified, show usage
	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filePath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filePath string) error {
	records, headers := readCsv(filePath)
	savePath := getSavePath(filePath)
	recMap := parseRecords(records, headers)
	jsonText := build(recMap)
	save(jsonText, savePath)
}

func readCsv(filePath string) ([][]string, []string) {
	// readCsv opens and reads a CSV file and returns:
	// * Records: a slice of slices of strings, one slice per row of the input
	// file, then one string per field in the row
	// * a slice of strings, each string is one field name from the header row
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	headerRow, err := r.Read()
	if err != nil {
		log.Fatal("Unable to parse input file as CSV for "+filePath, err)
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse input file as CSV for "+filePath, err)
	}

	return records, headerRow
}

func parseRecords(records [][]string, headers []string) []map[string]interface{} {
	// parseRecords accepts records and headers as string slices and returns a
	// JSON-ready slice of maps with interface values.  One map is created for
	// each row in the original data
	var result []map[string]interface{}
	for i := 0; i < len(records); i++ {
		row := make(map[string]interface{})
		for j := 0; j < len(headers); j++ {
			row[headers[j]] = records[i][j]
		}
		result = append(result, row)
	}
	return result
}

func build(m []map[string]inferface{}) []string {
	// build reconfigures maps into single-line JSON formatted data.
	// Each map in the slice is marshalled into a new JSON-formatted
	// string in the return slice.
	var result []string
	for i := 0; i < len(m); i++ {
		row := json.marshal(m[i])
		result = append(result, row)
	}
	return result
}

func getSavePath(inFile string) string {
	// getSavePath creates a JSON file name and path for the inputted CSV path.
	outFile := strings.Replace(inFile, ".csv", ".json", -1)
	return outFile
}

func save(text []string, filePath string) {
	// save saves the row-wise JSON data as a valid JSON file.
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for _, row := range text {
		fmt.Fprintln(f, row)
	}
	return
}