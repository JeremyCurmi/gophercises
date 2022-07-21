package main

import (
	"github.com/JeremyCurmi/gophercises/001-quiz-game/src/config"
	"github.com/JeremyCurmi/gophercises/001-quiz-game/src/csv"
	"github.com/JeremyCurmi/gophercises/001-quiz-game/src/logic"
)

func main() {
	data := csv.ReadCSVFile(*config.Filename)
	if *config.ShuffleQuestions {
		data = csv.ShuffleData(data)
	}
	logic.StartQuiz(data, *config.TimeLimit)
}
