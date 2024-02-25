package test1

import (
	"fmt"
)

// Define an interface named Writer
type Writer interface {
	Write([]byte) (int, error)
}

// Define a struct named File that implements the Writer interface
type File struct {
	// Some file-specific fields
}

// Implement the Write method for the File struct
func (f *File) Write(data []byte) (int, error) {
	// Simulate writing data to a file
	fmt.Println("Writing data to file:", string(data))
	return len(data), nil
}

// Define a function that takes a Writer interface as a parameter
func WriteToInterface(w Writer, data []byte) error {
	// Call the Write method of the Writer interface
	_, err := w.Write(data)
	return err
}

func RunTest1() {
	// Create a new File struct
	file := &File{}

	// Call the writeToInterface function with the File struct
	err := WriteToInterface(file, []byte("Hello, world!"))
	if err != nil {
		fmt.Printf("Error writing data: %v\n", err)
		return
	}

	fmt.Println("Data written successfully!")
}
