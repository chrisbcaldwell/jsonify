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
)

// Intermediate data structure "constants"
var goldenHeaders = []string{"name", "age", "DOB", "tacos", "burritos"}
var goldenRecords = [][]string{
	{"Chris", "47", "04-09-1978", "35", "24"},
	{"Carl", "23", "12-12-2002", "3", "1"},
	{"Carla", "7", "03-05-2018", "6", "3"},
}
var goldenMap = []map[string]interface{}{
	{"name": "Chris", "age": "47", "DOB": "04-09-1978", "tacos": "35", "burritos": "24"},
	{"name": "Carl", "age": "23", "DOB": "12-12-2002", "tacos": "3", "burritos": "1"},
	{"name": "Carla", "age": "7", "DOB": "03-05-2018", "tacos": "6", "burritos": "3"},
}
var goldenJson = []string{
	"{{\"name\":\"Chris\",\"age\":\"47\",\"DOB\":\"04-09-1978\",\"tacos\":\"35\",\"burritos\":\"24\"}}",
	"{{\"name\":\"Carl\",\"age\":\"23\",\"DOB\":\"12-12-2002\",\"tacos\":\"3\",\"burritos\":\"1\"}}",
	"{{\"name\":\"Carla\",\"age\":\"7\",\"DOB\":\"03-05-2018\",\"tacos\":\"6\",\"burritos\":\"3\"}}",
}

func TestReadCsv(t *testing.T) {
	_, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	testRecords, testHeaders, err := readCsv(inputFile)
	if !reflect.DeepEqual(goldenHeaders, testHeaders) {
		t.Logf("expected headers:\n%v", goldenHeaders)
		t.Logf("result headers:\n%v", testHeaders)
		t.Error("Result headers do not match expected")
	}
	if !reflect.DeepEqual(goldenRecords, testRecords) {
		t.Logf("expected records:\n%v", goldenRecords)
		t.Logf("result records:\n%v", testRecords)
		t.Error("Result records do not match expected")
	}
	if err != nil {
		t.Error("Error extracting csv data")
	}

}

func TestParseRecords(t *testing.T) {
	testMap, err := parseRecords(goldenRecords, goldenHeaders)
	if !reflect.DeepEqual(goldenMap, testMap) {
		t.Logf("expected map:\n%v", goldenMap)
		t.Logf("result map:\n%v", testMap)
		t.Error("Result map does not match expected")
	}
	if err != nil {
		t.Error("Error compiling data into map structure")
	}
}

func TestBuild(t *testing.T) {
	testJson, err := build(goldenMap)
	if !reflect.DeepEqual(goldenJson, testJson) {
		t.Logf("expected slice of JSON lines:\n%v", goldenJson)
		t.Logf("result slice of JSON lines:\n%v", testJson)
		t.Error("Result JSON lines slice does not match expected")
	}
	if err != nil {
		t.Error("Error compiling JSON lines")
	}

}
