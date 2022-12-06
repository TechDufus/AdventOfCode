package helpers

import (
	"bufio"
	"log"
	"os"
	"time"
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

// StartTimer starts a timer and returns the time
func StartTimer() time.Time {
	return time.Now()
}

// GetRuntimeMs() returns the time since the timer was started
func GetRuntimeMs(start time.Time) time.Duration {
	return time.Duration(time.Since(start).Microseconds())
}
