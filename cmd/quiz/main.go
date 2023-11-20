package main

import (
	"github.com/lorenz/quiz/internal/setup"
	"github.com/lorenz/quiz/internal/word"
)

var workSet word.Set

func main() {
	setup.DisplayMenu()
	workSet = word.CreatSet(setup.ReadIn())
}