package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run . 'input'")
		return
	}
	ascii, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Read Error.")
		return
	}

	lines := strings.Split(string(ascii), "\n")
	asciiarr := make([][]string, 0)
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		asciiarr = append(asciiarr, lines[i:end])
	}

	input := os.Args[1]
	input = strings.ReplaceAll(input, "\\n", "\n")

	output := make([]string, 9)
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			output = append(output, "\n")
		} else {
			for j, line := range asciiarr[input[i]-32] {
				output[j] += line
			}
		}
	}

	for _, line := range output {
		fmt.Println(line)
	}
}
