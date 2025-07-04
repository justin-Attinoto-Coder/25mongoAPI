package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Print welcome message
    welcome := "Welcome to user input"
    fmt.Println(welcome)

    // Create a scanner to read user input
    scanner := bufio.NewScanner(os.Stdin)

    // Prompt for name
    fmt.Print("Enter your name: ")
    scanner.Scan()
    name := strings.TrimSpace(scanner.Text())

    // Prompt for age
    fmt.Print("Enter your age: ")
    scanner.Scan()
    ageInput := scanner.Text()

    // Store inputs in an interface slice to demonstrate comma ok syntax
    inputs := []interface{}{name, ageInput}

    // Use comma ok syntax for type assertion
    for i, input := range inputs {
        if value, ok := input.(string); ok {
            fmt.Printf("Input %d is a string: %s\n", i+1, value)
        } else {
            fmt.Printf("Input %d is not a string\n", i+1)
        }
    }

    // Read a line to demonstrate bufio.Scanner for bytes and lines
    fmt.Print("Enter a short bio: ")
    scanner.Scan()
    bio := scanner.Text()
    fmt.Printf("Your bio: %s\n", bio)

    // Count runes in the bio to demonstrate rune scanning
    runeCount := len([]rune(bio))
    fmt.Printf("Your bio has %d runes\n", runeCount)

    // Example of reading a single word
    fmt.Print("Enter a favorite word: ")
    scanner.Scan()
    word := scanner.Text()
    fmt.Printf("Your favorite word: %s\n", word)
}