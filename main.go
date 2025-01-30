package main

import (
	"ascii-art/function"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run . >---> input")
		return
	}
	ascii, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Read Error.")
		return
	}
	input := os.Args[1]
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("inprintable char not allowed")
			return
		}
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	lines := function.Split(string(ascii))
	asciiarr := make(map[rune][]string)
	chatr := 32
	for _, v := range lines {
		asciiarr[rune(chatr)] = v
		chatr++

	}

	var slice []string // slice
	str := ""          // str
	invalid := false   // isfinish ola islastelement

	if input[len(input)-1] == '\n' {
		invalid = true
	}

	for _, x := range input {
		if x != '\n' {
			str += string(x)
		} else {
			if str != "" {
				slice = append(slice, str)
				str = ""
			}
			if x == '\n' {
				slice = append(slice, "\n")
			}

		}
	}
	if str != "" {
		slice = append(slice, str)
		if invalid {
			slice = append(slice, "\n")
		}
	}

	counter := 0 // counter
	for _, n := range slice {
		if n == "\n" {
			counter++
		}

	}
	if counter == len(slice) {
		for counter > 0 {
			fmt.Println()
			counter--

		}
		return
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != "\n" {
			lastElement := function.Rida(slice[i], asciiarr)
			for i := 0; i < len(lastElement); i++ {
				fmt.Println(lastElement[i])

			}
		} else {
			if i > 1 && i < len(slice)-1 {
				if !(slice[i-1] != "\n" && slice[i+1] != "\n") {
					fmt.Println()
				}
			}
		}

	}

}
