package helpers

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestGetInput(t *testing.T) {
	const testFile = "test.txt"

	// Create a test file with some sample text
	f, err := os.Create(testFile)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString("test\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.WriteString("123\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.WriteString("hello world\n")
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

	// Test the GetInput function
	lines := GetInput(testFile)
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "test" {
		t.Errorf("Expected first line to be 'test', got '%s'", lines[0])
	}
	if lines[1] != "123" {
		t.Errorf("Expected second line to be '123', got '%s'", lines[1])
	}
	if lines[2] != "hello world" {
		t.Errorf("Expected third line to be 'hello world', got '%s'", lines[2])
	}

	// Clean up the test file
	os.Remove(testFile)
}
