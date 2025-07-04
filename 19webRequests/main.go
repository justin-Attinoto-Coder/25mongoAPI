package main

import (
    "fmt"
    "io"
    "net/http"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our web adventure
    fmt.Println("Welcome to our fruit shop web request adventure!")

    // Example 1: Fetch data from lco.dev
    const url = "https://lco.dev"
    fmt.Println("\nFetching from", url)
    fetchWebsite(url)

    // Example 2: Fetch data from another site
    const fruitUrl = "https://example.com" // Pretend it's a fruit site
    fmt.Println("\nFetching from", fruitUrl)
    fetchWebsite(fruitUrl)

    // Example 3: Start a fruit shop web server
    fmt.Println("\nStarting fruit shop web server on http://localhost:8080")
    http.HandleFunc("/", serveFruitShop)      // Home page
    http.HandleFunc("/fruits", listFruits)    // Fruit list page
    http.ListenAndServe(":8080", nil)         // Start server
}

// Helper to fetch data from a website, like asking for a toy
func fetchWebsite(url string) {
    // Send a request to the website
    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    // Always close the response, like shutting a toy box
    defer response.Body.Close() // Hitesh says: Caller's job to close!

    // Check what kind of toy box we got
    fmt.Printf("Response is of type: %T\n", response)

    // Read all the toys (data) from the box
    dataBytes, err := io.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // Turn the data into words we can read
    content := string(dataBytes)
    // Show just a bit of the content (not too much!)
    if len(content) > 100 {
        fmt.Println("First 100 letters of content:\n", content[:100], "...")
    } else {
        fmt.Println("Content:\n", content)
    }
}

// Helper to serve the fruit shop home page
func serveFruitShop(w http.ResponseWriter, r *http.Request) {
    // io.ReadWriter doesn't close responses, so we don't need to close w
    fmt.Fprintf(w, "<h1>Welcome to the Fruit Shop!</h1><p>Visit <a href='/fruits'>Fruits</a> to see our stock!</p>")
}

// Helper to serve a list of fruits
func listFruits(w http.ResponseWriter, r *http.Request) {
    fruits := []string{"Apple", "Banana", "Mango", "Orange"}
    fmt.Fprintf(w, "<h1>Our Fruits</h1><ul>")
    for _, fruit := range fruits {
        fmt.Fprintf(w, "<li>%s</li>", fruit)
    }
    fmt.Fprintf(w, "</ul>")
}

// Explanation comments for a 5-year-old:
// - What's a web request? It's like asking a website for a toy (like a page) or sharing your toys with others!
// - What's 'http.Get(url)'? It's like sending a letter to a website (like lco.dev) asking for its toys.
// - What's 'response.Body.Close()'? It's like shutting a toy box after you take out the toys, so it stays tidy. Hitesh says we must do this!
// - What's 'defer response.Body.Close()'? It's like promising to shut the box when you're done, even if you forget later.
// - What's 'io.ReadAll(response.Body)'? It's like taking all the toys (words) out of the website's box to look at them.
// - What's 'string(dataBytes)'? It's like turning the toys into words we can read, like a story.
// - What's 'http.HandleFunc("/", serveFruitShop)'? It's like setting up a shop stall where people can visit and see your toys (pages).
// - What's 'serveFruitShop(w, r)'? It's like giving visitors a welcome sign for your fruit shop when they come.
// - What's 'fmt.Fprintf(w, ...)'? It's like writing words on a sign for visitors to read, like "Welcome to the Fruit Shop!"
// - Why doesn't 'w' need closing? Hitesh says io.ReadWriter (like 'w') doesn't close responses; the server handles it for us.
// - What's 'http.ListenAndServe(":8080", nil)'? It's like opening your fruit shop at a special address (localhost:8080) for everyone to visit.
// - Why fruit shop? It keeps the fun fruit theme (like Apple, Mango) from earlier chapters, like running a shop on the web.
// - How does this help beginners? It shows how to ask websites for toys (client) and share your own toys (server) with simple examples, plus reminds you to close toy boxes (responses) to be a good coder.