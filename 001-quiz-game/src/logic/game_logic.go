package logic

import (
	"fmt"
	"github.com/JeremyCurmi/gophercises/001-quiz-game/src/input"
	"time"
)

func StartQuizPart1(data [][]string) {
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

func StartQuiz(data [][]string, timeLimit int) {
	nQuestions, correctAnswers := len(data), 0

	ticker := time.NewTicker(time.Second * time.Duration(timeLimit))
	done := make(chan bool)

	go func() {
		for i := 0; i < nQuestions; i++ {
			question := data[i][0]
			solution := data[i][1]
			answer := input.ReadUserInput(question + ": ")
			if solution == answer {
				correctAnswers++
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-ticker.C:
		fmt.Println("Time's Up!")
	}
	fmt.Println("correctAnswers: ", correctAnswers, "out of: ", nQuestions)

}
