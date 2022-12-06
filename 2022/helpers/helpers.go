package helpers

import (
	"bufio"
	"log"
	"os"
)

// GetInput returns the input file as a slice of strings
func GetInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
