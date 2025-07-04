package main

import "fmt"

func main() {
    // Print a welcome message
    fmt.Println("Welcome to Go Programming!")

    // Declare and initialize variables
    name := "Justin"
    age := 41

    // Print variables with formatting
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

    // Basic arithmetic
    yearsLater := age + 5
    fmt.Printf("In 5 years, %s will be %d years old.\n", name, yearsLater)
}