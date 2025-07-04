package main

import "fmt"

// A user toy box holds info about a person
type user struct {
    name string // User's name, like "Justin"
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our defer adventure
    fmt.Println("Welcome to our fruit shop cleanup adventure with defer!")

    // Example 1: Simple defer in fruit shop
    fmt.Println("\nExample 1: Cleaning up the fruit shop")
    sellFruits()
    // Defers cleanup until sellFruits is done

    // Example 2: Multiple defers for user session
    fmt.Println("\nExample 2: User session with logging")
    justin := user{name: "Justin"}
    manageSession(justin)
    // Defers logout and log saving

    // Example 3: Simulated file operation
    fmt.Println("\nExample 3: Writing to a fruit order file")
    writeFruitOrder()
    // Defers file closing
}

// Helper to sell fruits in the shop
func sellFruits() {
    // Save cleanup for later, like promising to tidy up after playing
    defer cleanupShop()
    fmt.Println("Selling apples and bananas...")
    fmt.Println("Checking fruit stock...")
    // Cleanup happens when we're done
}

// Helper to clean up the fruit shop
func cleanupShop() {
    fmt.Println("Sweeping the shop floor!")
    fmt.Println("Putting away fruit baskets!")
}

// Helper to manage a user session
func manageSession(u user) {
    // Save these tasks for later, like promising to finish chores
    defer fmt.Println("Logging out", u.name) // Last to run
    defer saveSessionLog(u.name)             // Runs before logout
    fmt.Println("Starting session for", u.name)
    fmt.Println("User browsing fruit catalog...")
    // Defers run in reverse order when done
}

// Helper to save session log
func saveSessionLog(name string) {
    fmt.Println("Saving session log for", name)
}

// Helper to simulate writing a fruit order to a file
func writeFruitOrder() {
    // Save closing the file for later
    defer closeFile()
    fmt.Println("Opening fruit order file...")
    fmt.Println("Writing order: 5 apples, 3 mangos")
    // Closing happens when we're done
}

// Helper to simulate closing a file
func closeFile() {
    fmt.Println("Closing fruit order file!")
}

// Explanation comments for a 5-year-old:
// - What's defer? It's like promising to do a chore (like cleaning) after you're done playing, so you don't forget!
// - What's 'defer cleanupShop()'? It's like saying, "I'll sweep the shop when I'm done selling fruits!"
// - What's 'sellFruits()'? It's like a job where you sell fruits, and defer makes sure cleanup happens at the end.
// - What's 'defer fmt.Println("Logging out", u.name)'? It's like promising to say "goodbye" to Justin when his shopping is done.
// - What's 'defer saveSessionLog(u.name)'? It's like promising to save a note about what Justin did in the shop.
// - Why multiple defers in 'manageSession'? They stack up like chores, and Go does them backward (last chore first) when the job is done.
// - What's 'writeFruitOrder()'? It's like writing a list of fruits to order, and defer makes sure we close the list when done.
// - What's 'closeFile()'? It's like shutting a notebook after writing in it.
// - Why use defer? It makes sure important tasks (like cleaning or closing) happen even if you forget or something goes wrong.
// - Why fruit shop? It keeps the fun fruit theme (like apples, mangos) from earlier chapters, like playing in a shop.
// - How does this help beginners? It shows how defer saves tasks for later with simple examples like cleaning a shop, logging out, or closing a file, making it easy to learn.