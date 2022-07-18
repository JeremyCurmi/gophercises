package config

import (
	"flag"
)

var (
	Filename = flag.String("FILENAME", "problems.csv", "quiz csv file path")
)
