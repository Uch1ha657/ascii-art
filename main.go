package main

import (
	"ascii-art/function"
	"fmt"
	"os"
	"strings"
)

func main() {
	currhash, err := function.GetHash("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	validhash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	if currhash != validhash {
		fmt.Println("invalid file hash, file is not the valid one.")
		return
	}
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
	// get every 8 lines
	lines := function.Split(string(ascii))
	// turn our lines to a map represented by runes
	asciiarr := make(map[rune][]string)
	chatr := 32
	for _, v := range lines {
		asciiarr[rune(chatr)] = v
		chatr++
	}

	// split our input to words and \n
	var slice []string
	str := ""
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
				str = ""
			}
		}
	}
	if str != "" {
		slice = append(slice, str)
		str = ""
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != "\n" {
			lastElement := function.GetArr(slice[i], asciiarr)
			for i := 0; i < len(lastElement); i++ {
				fmt.Println(lastElement[i])
			}
		} else {
			if i+1 < len(slice) {
				if slice[i] == "\n" && slice[i+1] == "\n" {
					fmt.Println()
				}
			}
		}
	}
}
