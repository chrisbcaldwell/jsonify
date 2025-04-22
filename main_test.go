package main

import (
	"os"
	"reflect"
	"testing"
)

// File constants
const (
	inputFile  = "./testdata/test1.csv"
	resultFile = "test1.json"
	goldenFile = "./testdata/test1.json"
	goldenJson = `{"name":"Chris","age":"47","DOB":"04-09-1978","tacos":"35","burritos":"24"}
{"name":"Carl","age":"23","DOB":"12-12-2002","tacos":"3","burritos":"1"}
{"name":"Carla","age":"7","DOB":"03-05-2018","tacos":"6","burritos":"3"}`
)

// Intermediate step "constants"
var goldenHeaders = [5]string{"name", "age", "DOB", "tacos", "burritos"}
var goldenRecords = [3][5]string{
	{"Chris", "47", "04-09-1978", "35", "24"},
	{"Carl", "23", "12-12-2002", "3", "1"},
	{"Carla", "7", "03-05-2018", "6", "3"},
}

func TestReadCsv(t *testing.T) {
	_, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	testRecords, testHeaders, err := readCsv(inputFile)
	if !reflect.DeepEqual(goldenHeaders[:], testHeaders) {
		t.Logf("expected headers:\n%v", csvHeaders)
		t.Logf("result headers:\n%v", testHeaders)
		t.Error("Result headers do not match expected")
	}

}
