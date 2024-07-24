package main

/*
   This is a simple Hello World program in Go.
   It prints "Hello, World!" to the console.
*/

// Import the fmt package for formatted I/O operations
import "fmt"

// Define a struct to represent a person
type Person struct {
	Name string
	Age  int
}

// Define a method for the Person struct
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Define an interface to represent a greeter
type Greeter interface {
	Greet() string
}

// Make the Person struct implement the Greeter interface
func (p Person) SayHello() {
	fmt.Println(p.Greet())
}

func main() {
	// Create a new Person instance
	p := Person{Name: "Alice", Age: 30}

	// Call the SayHello method
	p.SayHello() // Output: Hello, my name is Alice
}
