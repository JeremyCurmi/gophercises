package config

import (
	"flag"
)

var (
	Filename         = flag.String("FILENAME", "problems.csv", "quiz csv file path")
	TimeLimit        = flag.Int("TIME_LIMIT", 10, "time limit to answer all quiz questions")
	ShuffleQuestions = flag.Bool("SHUFFLE_QUESTIONS", false, "Shuffle quiz questions")
)
