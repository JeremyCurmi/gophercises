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
			return strings.TrimSuffix(input, "\n")
		}
		log.Printf("Error reading user input: %v. Please try again.", err)
	}
}
