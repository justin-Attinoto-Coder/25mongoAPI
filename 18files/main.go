package main

import (
    "fmt"
    "io"
    "os"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our file adventure
    fmt.Println("Welcome to files in golang!")

    // Create a file to store a fruit order, like making a new notebook
    content := "Fruit order: 5 Apples, 3 Bananas - FruitShop.in"
    file, err := os.Create("./myfruitfile.txt")
    checkNilErr(err) // Check if something went wrong

    // Write the fruit order to the file, like writing in the notebook
    length, err := io.WriteString(file, content)
    checkNilErr(err)
    fmt.Println("Wrote", length, "letters to the file!")
    file.Close() // Close the notebook when done

    // Read the file to see the order, like reading the notebook
    readFile("./myfruitfile.txt")

    // Add more fruits to the file, like adding to the notebook
    appendToFile("./myfruitfile.txt", "\nExtra: 2 Mangos")
    fmt.Println("Added more fruits to the file!")

    // Read the file again to see everything
    readFile("./myfruitfile.txt")
}

// Helper to read a file, like reading a notebook
func readFile(filename string) {
    data, err := os.ReadFile(filename) // Read all the words
    checkNilErr(err)
    fmt.Println("Text in the file is:\n", string(data))
}

// Helper to append text to a file, like adding more notes
func appendToFile(filename, content string) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644) // Open for adding
    checkNilErr(err)
    defer file.Close() // Promise to close the notebook later
    _, err = io.WriteString(file, content)
    checkNilErr(err)
}

// Helper to check if something went wrong, like spotting a broken toy
func checkNilErr(err error) {
    if err != nil {
        panic(err) // Stop if there's a problem
    }
}

// Explanation comments for a 5-year-old:
// - What's a file? It's like a notebook where you can write toys (words) and read them later!
// - What's 'os.Create("./myfruitfile.txt")'? It's like making a new notebook called "myfruitfile.txt" to write fruit orders.
// - What's 'io.WriteString(file, content)'? It's like writing words (like "5 Apples") in the notebook.
// - What's 'length, err := io.WriteString(...)'? It's like counting how many letters you wrote and checking if it worked.
// - What's 'file.Close()'? It's like shutting the notebook so it's safe.
// - What's 'os.ReadFile(filename)'? It's like opening the notebook and reading all the words inside.
// - What's 'string(data)'? It's like turning the notebook's words into something we can show on the screen.
// - What's 'appendToFile(...)'? It's like opening the notebook again and adding more words (like "2 Mangos") at the end.
// - What's 'os.OpenFile(..., os.O_APPEND|os.O_WRONLY, 0644)'? It's like opening the notebook to add words without erasing old ones.
// - What's 'defer file.Close()'? It's like promising to shut the notebook when you're done adding words.
// - What's 'checkNilErr(err)'? It's like a helper that checks if something broke (like a toy) and stops the game if it did.
// - Why fruit orders? It keeps the fruit theme (like Apples, Bananas) from earlier chapters, making it fun like a shop game.
// - How does this help beginners? It shows how to make, write, read, and add to notebooks (files) with simple fruit orders, and uses a helper to catch mistakes, making files easy to learn.