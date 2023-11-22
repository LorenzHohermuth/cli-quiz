package question

import (
    "fmt"
	w "github.com/lorenz/quiz/internal/word"
	"math/rand"
	"time"
)

type MultipleChoiceQuestion struct {
	sugg  []w.Word
	answer w.Word
}

func NewMultiQuestion(sugg []w.Word, answer w.Word) *MultipleChoiceQuestion {
	return &MultipleChoiceQuestion{sugg, answer}
}

func (m *MultipleChoiceQuestion) Ask() {
	words := [4]w.Word{m.sugg[0], m.sugg[1], m.sugg[2], m.answer}
	words = mix(words)
	fmt.Printf("d: %v	f: %v\n", words[0].GetTerm(), words[1].GetTerm())
	fmt.Printf("j: %v	k: %v\n", words[2].GetTerm(), words[3].GetTerm())
}

func mix(arr [4]w.Word) [4]w.Word {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func (m *MultipleChoiceQuestion) Check() bool {
    return true
}