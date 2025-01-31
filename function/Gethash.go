package function

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func GetHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()

	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
