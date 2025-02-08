package function

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func Split(ascii string) [][]string {
	slise := []string{}
	lines := [][]string{}
	line := ""
	for _, v := range ascii {
		if v == '\n' {
			slise = append(slise, line)
			line = ""

		} else {
			line += string(v)
		}
	}
	for i := 1; i < len(slise); i += 8 {
		if i+8 > len(slise) {
			break
		}
		abed := slise[i : i+8]
		lines = append(lines, abed)
		i++
	}
	return lines
}

// get our intended result
func GetArr(input string, asciiarr map[rune][]string) []string {
	ret := make([]string, 8)
	for _, v := range input {
		ris := asciiarr[v]
		for i := 0; i < 8; i++ {
			ret[i] = ret[i] + ris[i]
		}
	}
	return ret
}

func GetHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// sha256 = secure hash algorithm (256bit)
	hash := sha256.New()

	// reads the content of the file and fnefs lwe9t writes it to the hash
	// meaning we update the hash progressively as the file is being read
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// return as a hexadecimal string (cuz hash is a bye slice)
	// cuz hashes are usually represented in hexadecimal
	// the hash.sum func computes and returns the final hash
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
