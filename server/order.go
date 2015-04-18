package main

import (
	"sort"
)

type IndexWord struct {
	Count int64
	Word  string
}

type IndexChar struct {
	Count int64
	Char  rune
}

type ByWord func(p1, p2 *IndexWord) bool

func (by ByWord) Sort(words []IndexWord) {
	s := &wordSorter{
		words: words,
		by:    by,
	}
	sort.Sort(s)
}

type wordSorter struct {
	words []IndexWord
	by    func(p1, p2 *IndexWord) bool
}

func (s *wordSorter) Len() int {
	return len(s.words)
}

func (s *wordSorter) Swap(i, j int) {
	s.words[i], s.words[j] = s.words[j], s.words[i]
}

func (s *wordSorter) Less(i, j int) bool {
	return s.by(&s.words[i], &s.words[j])
}

type ByChar func(p1, p2 *IndexChar) bool

func (by ByChar) Sort(chars []IndexChar) {
	s := &charSorter{
		chars: chars,
		by:    by,
	}
	sort.Sort(s)
}

type charSorter struct {
	chars []IndexChar
	by    func(p1, p2 *IndexChar) bool
}

func (s *charSorter) Len() int {
	return len(s.chars)
}

func (s *charSorter) Swap(i, j int) {
	s.chars[i], s.chars[j] = s.chars[j], s.chars[i]
}

func (s *charSorter) Less(i, j int) bool {
	return s.by(&s.chars[i], &s.chars[j])
}
