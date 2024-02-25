package test1_test

import (
	"bytes"
	"testing"

	"github.com/taylormonacelli/hersheep/test1"
)

// TestWriteToInterface tests the writeToInterface function
func TestWriteToInterface(t *testing.T) {
	// Create a buffer to simulate a writer
	buffer := &bytes.Buffer{}
	// Call writeToInterface with the buffer
	err := test1.WriteToInterface(buffer, []byte("Testing writeToInterface"))
	// Check if error occurred
	if err != nil {
		t.Errorf("Error writing data: %v", err)
	}
	// Check if data was written to the buffer
	if buffer.String() != "Testing writeToInterface" {
		t.Errorf("Expected buffer to contain 'Testing writeToInterface', got '%s'", buffer.String())
	}
}
