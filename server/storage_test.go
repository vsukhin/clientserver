package main

import (
	"testing"
)

type TestLogger struct {
}

func (logger *TestLogger) Println(args ...interface{}) {
}

func (logger *TestLogger) Fatalf(query string, args ...interface{}) {
}

func TestNewResults(t *testing.T) {
	var total int64 = 1
	var topwords = []string{"1"}
	var topchars = []string{"1"}

	results := NewResults(total, topwords, topchars)
	if results == nil {
		t.Error("New results should create a new object")
	} else {
		if results.Total != total {
			t.Error("New results should use a right total value")
		}
		if len(results.TopWords) != 1 && results.TopWords[0] != "1" {
			t.Error("New results should use a right word slice")
		}
		if len(results.TopChars) != 1 && results.TopChars[0] != "1" {
			t.Error("New results should use a right char slice")
		}
	}
}

func TestNewStorage(t *testing.T) {
	var useRawText = true

	storage := NewStorage(useRawText)
	if storage == nil {
		t.Error("New storage should create a new object")
	} else {
		if storage.useRawText != useRawText {
			t.Error("New storage should use a right use raw text value")
		}
	}
}

func TestGetWordsTotal(t *testing.T) {
	Log = new(TestLogger)
	useRawText := false
	storage := NewStorage(useRawText)
	storage.WordIndex["A"] = 1
	storage.WordIndex["B"] = 2
	storage.WordIndex["C"] = 3

	total := storage.GetWordsTotal()
	if total != storage.WordIndex["A"]+storage.WordIndex["B"]+storage.WordIndex["C"] {
		t.Error("Get words total should calculate a proper value")
	}
}

func TestGetTopWords(t *testing.T) {
	Log = new(TestLogger)
	useRawText := false
	storage := NewStorage(useRawText)
	storage.WordIndex["A"] = 1
	storage.WordIndex["B"] = 2
	storage.WordIndex["C"] = 3
	storage.WordIndex["D"] = 4
	storage.WordIndex["E"] = 5
	storage.WordIndex["F"] = 6

	topwords := storage.GetTopWords()
	if len(topwords) != TOP_COUNT {
		t.Error("Get top words should return a proper number of words")
	} else {
		if topwords[0] != "F" && topwords[1] != "E" && topwords[2] != "D" && topwords[3] != "C" && topwords[4] != "B" {
			t.Error("Get top words should return proper ordering words")
		}
	}
}

func TestGetTopChars(t *testing.T) {
	Log = new(TestLogger)
	useRawText := false
	storage := NewStorage(useRawText)
	storage.CharIndex['A'] = 1
	storage.CharIndex['B'] = 2
	storage.CharIndex['C'] = 3
	storage.CharIndex['D'] = 4
	storage.CharIndex['E'] = 5
	storage.CharIndex['F'] = 6

	topchars := storage.GetTopChars()
	if len(topchars) != TOP_COUNT {
		t.Error("Get top chars should return a proper number of chars")
	} else {
		if topchars[0] != "F" && topchars[1] != "E" && topchars[2] != "D" && topchars[3] != "C" && topchars[4] != "B" {
			t.Error("Get top chars should return proper ordering chars")
		}
	}
}
