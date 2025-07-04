package main

import (
    "encoding/json"
    "fmt"
)

// A course toy box holds info about a course
type course struct {
    Name     string   `json:"coursename"` // Course name, like "ReactJS Bootcamp"
    Price    int      `json:"price"`      // Price, like 299
    Website  string   `json:"website"`    // Platform, like "FruitLearnOnline.in"
    Password string   `json:"-"`          // Password, hidden in JSON
    Tags     []string `json:"tags"`       // Tags, like ["web", "js"]
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our JSON adventure
    fmt.Println("Welcome to our JSON toy box adventure!")

    // Example 1: Encode courses to JSON
    EncodeJson()

    // Example 2: Decode JSON to courses
    DecodeJson()

    // Example 3: Consume JSON from the web with struct and map
    ConsumeWebJson()
}

// Helper to encode courses to JSON, like turning toy boxes into words
func EncodeJson() {
    // Make a list of course toy boxes
    lcoCourses := []course{
        {
            Name:     "ReactJS Bootcamp",
            Price:    299,
            Website:  "FruitLearnOnline.in",
            Password: "abc123",
            Tags:     []string{"web", "javascript"},
        },
        {
            Name:     "MERN Bootcamp",
            Price:    199,
            Website:  "FruitLearnOnline.in",
            Password: "bcd123",
            Tags:     []string{"mern", "fullstack"},
        },
        {
            Name:     "Angular Bootcamp",
            Price:    299,
            Website:  "FruitLearnOnline.in",
            Password: "hit123",
            Tags:     nil,
        },
    }

    // Turn the toy boxes into pretty JSON words
    jsonData, err := json.MarshalIndent(lcoCourses, "", "  ")
    if err != nil {
        panic(err)
    }

    // Show the JSON words
    fmt.Println("\nEncoded JSON:")
    fmt.Println(string(jsonData))
}

// Helper to decode JSON to courses, like turning words into toy boxes
func DecodeJson() {
    // Pretend we got these JSON words from a file
    jsonData := []byte(`
    [
        {
            "coursename": "Go Bootcamp",
            "price": 99,
            "website": "FruitLearnOnline.in",
            "tags": ["go", "backend"]
        },
        {
            "coursename": "Python Bootcamp",
            "price": 149,
            "website": "FruitLearnOnline.in",
            "tags": null
        }
    ]
    `)

    // Check if the JSON words are okay
    if !json.Valid(jsonData) {
        fmt.Println("\nOops, bad JSON words from file!")
        return
    }

    // Make a box for course toys
    var courses []course

    // Turn JSON words into course toy boxes
    err := json.Unmarshal(jsonData, &courses)
    if err != nil {
        panic(err)
    }

    // Show the course toys
    fmt.Println("\nDecoded courses from file:")
    for i, c := range courses {
        fmt.Printf("Course %d: Name=%s, Price=%d, Website=%s, Tags=%v\n",
            i+1, c.Name, c.Price, c.Website, c.Tags)
    }
}

// Helper to consume JSON from the web, like getting toys from a website
func ConsumeWebJson() {
    // Pretend we got these JSON words from a web shop
    jsonDataFromWeb := []byte(`
    {
        "coursename": "NodeJS Bootcamp",
        "price": 299,
        "website": "FruitLearnOnline.in",
        "extra": "Special Offer"
    }
    `)

    // Check if the web JSON words are okay
    checkValid := json.Valid(jsonDataFromWeb)
    if checkValid {
        fmt.Println("\nJSON from web was valid!")

        // Try with a course toy box
        var lcoCourse course
        err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Web course (struct): %#v\n", lcoCourse)

        // Try with a magic toy box (map) for any toys
        var myonlineData map[string]interface{}
        err = json.Unmarshal(jsonDataFromWeb, &myonlineData)
        if err != nil {
            panic(err)
        }
        fmt.Println("Web course (map):")
        for key, value := range myonlineData {
            fmt.Printf("  %s: %v\n", key, value)
        }
    } else {
        fmt.Println("\nOops, bad JSON from web!")
    }

    // Try another web JSON with different toys
    jsonDataExtra := []byte(`
    {
        "coursename": "Java Bootcamp",
        "price": 249,
        "website": "FruitLearnOnline.in",
        "tags": ["java", "backend"],
        "discount": 10
    }
    `)

    // Check and decode with map
    if json.Valid(jsonDataExtra) {
        var extraData map[string]interface{}
        err := json.Unmarshal(jsonDataExtra, &extraData)
        if err != nil {
            panic(err)
        }
        fmt.Println("\nExtra web course (map):")
        for key, value := range extraData {
            fmt.Printf("  %s: %v\n", key, value)
        }
    } else {
        fmt.Println("\nOops, bad extra JSON from web!")
    }
}

// Explanation comments for a 5-year-old:
// - What's JSON? It's like turning toy boxes (courses) into words to send to friends, then turning words back into toy boxes!
// - What's consuming JSON? It's like getting words from a website and turning them into toys we can play with.
// - What's 'type course struct'? It's like making a toy box called "course" to hold toys like name, price, and website.
// - What's 'json:"coursename"'? It's like giving a toy a special name ("coursename") for JSON words.
// - What's 'Password string `json:"-"`'? It's like hiding the password toy so it doesn't show in JSON.
// - What's 'lcoCourses := []course{...}'? It's like making a shelf with course toy boxes, like "ReactJS Bootcamp".
// - What's 'json.MarshalIndent(lcoCourses, ...)'? It's like turning course boxes into pretty JSON words with spaces.
// - What's 'json.Valid(jsonDataFromWeb)'? It's like checking if web words are good JSON before opening them.
// - What's 'json.Unmarshal(jsonDataFromWeb, &lcoCourse)'? It's like turning web JSON words into a course toy box.
// - What's 'var myData map[string]interface{}'? It's like a magic toy box that can hold any toys from JSON, like a surprise bag!
// - What's 'json.Unmarshal(jsonDataFromWeb, &myData)'? It's like putting web JSON words into the magic box to see all toys.
// - What's 'fmt.Printf("%#v\n", lcoCourse)'? It's like showing all toys in a course box with their names, like a toy list.
// - Why use a map? It's like a flexible box for when we don't know what toys (like "extra" or "discount") the JSON has.
// - Why two web JSONs? It shows how to get different toys from the web, with structs (fixed toys) and maps (any toys).
// - Why fruit platform? It keeps the fun fruit theme (like FruitLearnOnline.in) from earlier, like a fruit school.
// - How does this help beginners? It shows how to make, read, and check JSON words with course toys, using both fixed boxes (structs) and magic boxes (maps), so you learn to share and use toys in Go.