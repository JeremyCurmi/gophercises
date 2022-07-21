package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadUserInput(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == nil {
			cleanInput := cleanInput(input)
			return strings.TrimSuffix(cleanInput, "\n")
		}
		log.Printf("Error reading user input: %v. Please try again.", err)
	}
}

func cleanInput(text string) string {
	var input string
	input = strings.ToLower(text)
	input = strings.TrimSpace(input)
	return input
}
