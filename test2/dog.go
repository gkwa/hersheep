package test2

import (
	"fmt"
)

// Define an interface named FoodProvider
type FoodProvider interface {
	ProvideFood() string
}

// Define a struct named Dog that implements the FoodProvider interface
type Dog struct {
	Name string
}

// Implement the ProvideFood method for the Dog struct
func (d *Dog) ProvideFood() string {
	// Simulate a dog providing its favorite food
	return fmt.Sprintf("%s's favorite food is: bones", d.Name)
}

// Define a function that takes a FoodProvider interface as a parameter
func FeedDog(fp FoodProvider) {
	// Call the ProvideFood method of the FoodProvider interface
	food := fp.ProvideFood()
	fmt.Println(food)
}

func RunTest2() {
	// Create a new Dog struct
	dog := &Dog{Name: "Buddy"}

	// Call the feedDog function with the Dog struct
	FeedDog(dog)
}
