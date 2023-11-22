package word

import (
	"strings"
)

type Word struct {
	index int
	rating uint8 // rating how good the word is learned between 0 and 100
	term string
	definition string
}

func CreatWord(index int, term string) Word {
	keyValue := strings.Split(term, "=")
	return Word{index, 0, keyValue[0], keyValue[1]}
}

func (w *Word) GetTerm() string {
	return w.term
}