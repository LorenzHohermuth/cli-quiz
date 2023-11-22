package word

import (
	"strings"
)

type Set struct {
	terms []Word
}

func CreatSet(s string) (set Set) {
	strarr := strings.Split(s, ";")
	newlen := len(strarr) - 1
	strarr = strarr[:newlen]
	arr := make([]Word, newlen)
	for i, v := range strarr {
		arr[i] = CreatWord(i, v)
	}
	return Set{arr}
}

func (s Set) Get(i int) Word {
	return s.terms[i]
}
