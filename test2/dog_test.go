package test2_test

import (
	"testing"

	"github.com/taylormonacelli/hersheep/test2"
)

func TestDogProvideFood(t *testing.T) {
	// Create a new Dog struct
	dog := &test2.Dog{Name: "Max"}

	// Call the ProvideFood method of the Dog struct
	result := dog.ProvideFood()

	// Check if the result matches the expected output
	expected := "Max's favorite food is: bones"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestDogName(t *testing.T) {
	// Create a new Dog struct
	dog := &test2.Dog{Name: "Rocky"}

	// Check if the dog's name is initialized correctly
	if dog.Name != "Rocky" {
		t.Errorf("Expected name to be 'Rocky' but got %q", dog.Name)
	}
}
