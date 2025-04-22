package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
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
	records, headers, err := readCsv(filePath)
	if err != nil {
		return err
	}
	savePath := getSavePath(filePath)
	recMap, err := parseRecords(records, headers)
	if err != nil {
		return err
	}
	jsonText, err := build(recMap)
	if err != nil {
		return err
	}
	err = save(jsonText, savePath)
	if err != nil {
		return err
	}
	fmt.Println("JSON file saved at " + savePath)
	return nil
}

func readCsv(filePath string) ([][]string, []string, error) {
	// readCsv opens and reads a CSV file and returns:
	// * Records: a slice of slices of strings, one slice per row of the input
	// file, then one string per field in the row
	// * a slice of strings, each string is one field name from the header row
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file " + filePath)
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	headerRow, err := r.Read()
	if err != nil {
		fmt.Println("Unable to parse input file as CSV for " + filePath)
		return nil, nil, err
	}

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse input file as CSV for " + filePath)
		return nil, nil, err
	}

	return records, headerRow, nil
}

func parseRecords(records [][]string, headers []string) ([]map[string]interface{}, error) {
	// parseRecords accepts records and headers as string slices and returns a
	// JSON-ready slice of maps with interface values.  One map is created for
	// each row in the original data
	var result []map[string]interface{}
	for i := 0; i < len(records); i++ {
		row := make(map[string]interface{})
		for j := 0; j < len(headers); j++ {
			if len(headers) != len(records[i]) {
				message := "record length does not match header length in all rows"
				fmt.Println(message)
				return nil, errors.New(message)
			}
			row[headers[j]] = records[i][j]
		}
		result = append(result, row)
	}
	return result, nil
}

func build(m []map[string]interface{}) ([]string, error) {
	// build reconfigures maps into single-line JSON formatted data.
	// Each map in the slice is marshalled into a new JSON-formatted
	// string in the return slice.
	var result []string
	for i := 0; i < len(m); i++ {
		row, err := json.Marshal(m[i])
		if err != nil {
			fmt.Println("Unable to format data as JSON")
			return nil, err
		}
		result = append(result, string(row))
	}
	return result, nil
}

func getSavePath(inFile string) string {
	// getSavePath creates a JSON lines file name and path for the inputted CSV path.
	outFile := inFile + ".jsonl"
	return outFile
}

func save(text []string, filePath string) error {
	// save saves the row-wise JSON data as a valid JSON lines file.
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error saving file")
		return err
	}
	defer f.Close()
	for _, row := range text {
		fmt.Fprintln(f, row)
	}
	return nil
}
