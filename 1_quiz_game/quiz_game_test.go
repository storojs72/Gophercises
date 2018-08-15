package main

import (
	"testing"
	"os"
	"encoding/csv"
	"reflect"
)

func TestCalculateGameResults(t *testing.T){
	var err error
	f, err := os.Open("problems_example.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	answers := []string {"10", "10", "2", "11", "3", "14", "4", "5", "6", "5"}
	result, err := CalculateGameResults(lines, answers)
	if err != nil {
		t.Fatal(err)
	}
	expectedResult := 10
	if result != expectedResult {
		t.Fatal("Expected:", expectedResult, "Got:", result)
	}

	answers = []string {"10", "4", "2", "11", "3", "11", "4", "5", "6", "5"}
	result, err = CalculateGameResults(lines, answers)
	if err != nil {
		t.Fatal(err)
	}
	expectedResult = 8
	if result != expectedResult {
		t.Fatal("Expected:", expectedResult, "Got:", result)
	}
}

func TestShuffle(t *testing.T){
	var err error
	f, err := os.Open("problems_example.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	newLines := Shuffle(lines)

	if reflect.DeepEqual(lines, newLines){
		t.Fatal("no shuffling: \n", lines, "\n", newLines)
	}
}


