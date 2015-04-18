package main

import (
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	TOP_COUNT = 5
)

type Storage struct {
	RawText   []string
	WordIndex map[string]int64
	CharIndex map[rune]int64

	sync.RWMutex
	useRawText bool
}

type Statistics interface {
	GetWordsTotal() (total int64)
	GetTopWords() (topwords []string)
	GetTopChars() (topchars []string)
	GetResults() (results *Results)
}

type Results struct {
	Total    int64    `json:"count"`
	TopWords []string `json:"top_5_words"`
	TopChars []string `json:"top_5_letters"`
}

func NewResults(total int64, topwords []string, topchars []string) (results *Results) {
	results = new(Results)
	results.Total = total
	results.TopWords = topwords
	results.TopChars = topchars

	return results
}

func NewStorage(useRawText bool) (storage *Storage) {
	storage = new(Storage)
	storage.RawText = *new([]string)
	storage.WordIndex = make(map[string]int64)
	storage.CharIndex = make(map[rune]int64)
	storage.useRawText = useRawText

	return storage
}

func (storage *Storage) ProcessText(textChannel chan string) {
	separators, err := regexp.Compile("[.,!?:;\\s]+")
	if err != nil {
		Log.Fatalf("Can't compile regular expression %v", err)
	}
	for {
		rawText := <-textChannel
		Log.Println("Starting processing raw text ", rawText, time.Now())
		words := separators.Split(strings.ToLower(rawText), -1)

		storage.Lock()
		if storage.useRawText {
			storage.RawText = append(storage.RawText, rawText)
		}
		for _, word := range words {
			storage.WordIndex[word]++
			for _, char := range word {
				storage.CharIndex[char]++
			}
		}
		storage.Unlock()

		Log.Println("Finishing processing raw text ", rawText, time.Now())
	}
}

func (storage *Storage) GetWordsTotal() (total int64) {
	Log.Println("Starting calculating words total ", time.Now())
	storage.RLock()
	defer storage.RUnlock()
	total = 0
	for _, count := range storage.WordIndex {
		total += count
	}
	Log.Println("Finishing calculating words total ", time.Now())

	return total
}

func (storage *Storage) GetTopWords() (topwords []string) {
	Log.Println("Starting calculating top words ", time.Now())
	topwords = *new([]string)
	allwords := *new([]IndexWord)
	storage.RLock()
	for key, value := range storage.WordIndex {
		allwords = append(allwords, IndexWord{Count: value, Word: key})
	}
	storage.RUnlock()

	count := func(w1, w2 *IndexWord) bool {
		return w1.Count > w2.Count
	}
	ByWord(count).Sort(allwords)

	for i := 0; i < len(allwords); i++ {
		if i < TOP_COUNT {
			topwords = append(topwords, allwords[i].Word)
		} else {
			break
		}
	}

	Log.Println("Finishing calculating top words ", time.Now())
	return topwords
}

func (storage *Storage) GetTopChars() (topchars []string) {
	Log.Println("Starting calculating top chars ", time.Now())
	topchars = *new([]string)
	allchars := *new([]IndexChar)
	storage.RLock()
	for key, value := range storage.CharIndex {
		allchars = append(allchars, IndexChar{Count: value, Char: key})
	}
	storage.RUnlock()

	count := func(w1, w2 *IndexChar) bool {
		return w1.Count > w2.Count
	}
	ByChar(count).Sort(allchars)

	for i := 0; i < len(allchars); i++ {
		if i < TOP_COUNT {
			topchars = append(topchars, string(allchars[i].Char))
		} else {
			break
		}
	}

	Log.Println("Finishing calculating top chars ", time.Now())
	return topchars
}

func (storage *Storage) GetResults() (results *Results) {

	return NewResults(storage.GetWordsTotal(), storage.GetTopWords(), storage.GetTopChars())
}
