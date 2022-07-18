package logic

import (
	"fmt"
	"github.com/JeremyCurmi/gophercises/001-quiz-game/src/input"
)

func StartQuiz(data [][]string) {
	nQuestions := len(data)
	correctAnswers := 0
	for i := 0; i < nQuestions; i++ {
		question := data[i][0]
		solution := data[i][1]
		answer := input.ReadUserInput(question + ": ")
		if solution == answer {
			correctAnswers++
		}
	}
	fmt.Println("correctAnswers: ", correctAnswers, "out of: ", nQuestions)
}
