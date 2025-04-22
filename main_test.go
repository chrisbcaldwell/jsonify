package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

// File constants
const (
	inputFile  = "./testdata/test1.csv"
	resultFile = "./testdata/test1.csv.jsonl"
	goldenFile = "./testdata/test1.jsonl"
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
	"{\"name\":\"Chris\",\"age\":\"47\",\"DOB\":\"04-09-1978\",\"tacos\":\"35\",\"burritos\":\"24\"}",
	"{\"name\":\"Carl\",\"age\":\"23\",\"DOB\":\"12-12-2002\",\"tacos\":\"3\",\"burritos\":\"1\"}",
	"{\"name\":\"Carla\",\"age\":\"7\",\"DOB\":\"03-05-2018\",\"tacos\":\"6\",\"burritos\":\"3\"}",
}

type jsonlMap []map[string]interface{}

var mapGolden jsonlMap
var mapTest jsonlMap

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
	if err != nil {
		t.Error("Error compiling map data into JSONL structure")
	}
	goldenBytes := []byte("[" + strings.Join(goldenJson, ",\n") + "]")
	testBytes := []byte("[" + strings.Join(testJson, ",\n") + "]")
	err = json.Unmarshal(goldenBytes, &mapGolden)
	if err != nil {
		t.Error("Error unmarshalling expected JSON data")
	}
	err = json.Unmarshal(testBytes, &mapTest)
	if err != nil {
		t.Error("Error unmarshalling result JSON data")
	}
	if !reflect.DeepEqual(mapGolden, mapTest) {
		t.Logf("expected slice of JSON lines:\n%v", goldenJson)
		t.Logf("result slice of JSON lines:\n%v", testJson)
		t.Error("Result JSON lines are not equivalent to expected")
	}
}

func TestRun(t *testing.T) {
	err := run(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(jsonlToJson(result), &mapTest)
	if err != nil {
		t.Error("Error unmarshalling result file")
	}
	err = json.Unmarshal(jsonlToJson(expected), &mapGolden)
	if err != nil {
		t.Error("Error unmarshalling expected file")
	}
	if !reflect.DeepEqual(mapTest, mapGolden) {
		t.Logf("expected:\n%s", expected)
		t.Logf("result:\n%s", result)
		t.Error("Result file does not match expected")
	}
	os.Remove(resultFile)
}
func jsonlToJson(jsonl []byte) []byte {
	s := strings.Replace(string(jsonl), "\n", ",\n", -1)
	s = "[" + s + "]"
	return []byte(s)
}
