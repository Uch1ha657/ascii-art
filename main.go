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

	input = strings.ReplaceAll(input, "\\n", "\n")
	lines := function.Split(string(ascii))
	asciiarr := make(map[rune][]string)
	chatr := 32
	for _, v := range lines {
		asciiarr[rune(chatr)] = v
		chatr++

	}
	
		var rida []string
		rid := ""
		ste := false

		if input[len(input)-1] == '\n'{
			ste = true
		}
	
		for _, x := range input{
			if x != '\n'{
				rid += string(x)
			}else{
				if rid!= ""{
					rida = append(rida, rid)
					rid = ""
				}
				if x == '\n'{
					rida = append(rida, "\n")
				}
				
			}
		}
		if rid != ""{
			rida = append(rida, rid)
			if ste{
				rida = append(rida,"\n")
			}
		}
		
	
		
	
	contr := 0
	for _, n := range rida {
		if n == "\n" {
			contr++
		}

	}
	if contr == len(rida) {
		for contr > 0 {
			fmt.Println()
			contr--

		}
		return
	}

	for i := 0; i < len(rida); i++ {
		if rida[i] != "\n" {
			nrida := function.Rida(rida[i],asciiarr)
			for i := 0; i < len(nrida); i++ {
				fmt.Println(nrida[i])

			}
		} else {
			if i > 1 && i < len(rida)-1 {
				if !(rida[i-1] != "\n" && rida[i+1] != "\n") {
					fmt.Println()
				}
			}
		}

	}

}
