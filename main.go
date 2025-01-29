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
	input := os.Args[1]

	input = strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(string(ascii), "\n")
	asciiarr := make([][]string, 0)

	for i := 1; i < len(lines); i += 9 {
		end := i + 8
		if end > len(lines) {
			end = len(lines)
		}
		asciiarr = append(asciiarr, lines[i:end])
	}

	currentLines := [][]string{}
	line := make([]string, 8)

	for i := 0; i < len(input); i++ {
		if input[i] != '\n' {
			for j := 0; j < 8; j++ {
				line[j] += asciiarr[input[i]-32][j]
			}
		} else {
			currentLines = append(currentLines, line)
			line = make([]string, 8)
		}
	}

	currentLines = append(currentLines, line)

	for _, asciiLine := range currentLines {
		for _, row := range asciiLine {
			fmt.Println(row)
		}
	}
}
