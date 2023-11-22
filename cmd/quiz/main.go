package main

import (
	s "github.com/lorenz/quiz/internal/setup"
	w "github.com/lorenz/quiz/internal/word"
	q "github.com/lorenz/quiz/internal/question"
)

var workSet w.Set

func main() {
	s.DisplayMenu()
	workSet = w.CreatSet(s.ReadIn())
	wordArr := []w.Word{workSet.Get(0), workSet.Get(1), workSet.Get(2)}
	quest := q.NewMultiQuestion(wordArr, workSet.Get(3))
	quest.Ask()
}