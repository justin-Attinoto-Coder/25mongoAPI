package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
)

// A user toy box holds data sent to the server
type user struct {
    Name string `json:"name"` // User's name, like "Hitesh"
    Age  int    `json:"age"`  // User's age, like 4
}

// A course toy box holds course data
type course struct {
    CourseName string `json:"coursename"` // Course name, like "Let's go with golang"
    Price      int    `json:"price"`      // Price, like 0
    Platform   string `json:"platform"`   // Platform, like "learncodeonline.in"
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our web server and request adventure
    fmt.Println("Welcome to our fruit shop web server adventure!")

    // Example 1: Make GET, JSON POST, and Form POST requests
    fmt.Println("\nMaking GET and POST requests")
    PerformGetRequest("http://localhost:8000/get")     // Needs server running
    PerformPostJsonRequest("http://localhost:8000/post")
    PerformPostFormRequest("http://localhost:8000/postform")
    PerformGetRequest("https://example.com") // Public site

    // Example 2: Start the fruit shop web server
    fmt.Println("\nStarting fruit shop web server on http://localhost:8000")
    http.HandleFunc("/", welcomeHandler)          // Home page
    http.HandleFunc("/get", getHandler)           // Get message page
    http.HandleFunc("/post", postHandler)         // JSON/form post page
    http.HandleFunc("/postform", postFormHandler) // Form post page
    err := http.ListenAndServe(":8000", nil)      // Open shop
    if err != nil {
        panic(err)
    }
}

// Helper to make a GET request, like asking for a toy
func PerformGetRequest(myurl string) {
    fmt.Println("Fetching GET from", myurl)

    // Send a letter to the website
    response, err := http.Get(myurl)
    if err != nil {
        fmt.Println("Oops, couldn't reach", myurl, ":", err)
        return
    }
    // Promise to shut the toy box
    defer response.Body.Close()

    // Check how the website replied
    fmt.Println("Status code:", response.StatusCode)
    fmt.Println("Content length:", response.ContentLength, "letters")

    // Grab all the toys (words)
    dataBytes, err := io.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // Turn toys into words
    content := string(dataBytes)
    // Show a bit if it's long
    if len(content) > 100 {
        fmt.Println("First 100 letters:\n", content[:100], "...")
    } else {
        fmt.Println("Content:\n", content)
    }
}

// Helper to make a JSON POST request, like sending a toy box
func PerformPostJsonRequest(myurl string) {
    fmt.Println("Sending JSON POST to", myurl)

    // Make a toy box with course info
    requestBody := strings.NewReader(`
    {
        "coursename": "Let's go with golang",
        "price": 0,
        "platform": "learncodeonline.in"
    }
    `)

    // Send the toy box
    response, err := http.Post(myurl, "application/json", requestBody)
    if err != nil {
        fmt.Println("Oops, couldn't send to", myurl, ":", err)
        return
    }
    // Promise to shut the toy box
    defer response.Body.Close()

    // Check the reply
    fmt.Println("Status code:", response.StatusCode)
    fmt.Println("Content length:", response.ContentLength, "letters")

    // Grab the reply toys
    dataBytes, err := io.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // Show the reply
    content := string(dataBytes)
    fmt.Println("Reply content:\n", content)
}

// Helper to make a Form POST request, like sending a note
func PerformPostFormRequest(myurl string) {
    fmt.Println("Sending Form POST to", myurl)

    // Make a note with fruit order info
    formData := url.Values{
        "fruit":    {"Mango"},
        "quantity": {"3"},
        "customer": {"Justin"},
    }

    // Send the note
    response, err := http.PostForm(myurl, formData)
    if err != nil {
        fmt.Println("Oops, couldn't send to", myurl, ":", err)
        return
    }
    // Promise to shut the toy box
    defer response.Body.Close()

    // Check the reply
    fmt.Println("Status code:", response.StatusCode)
    fmt.Println("Content length:", response.ContentLength, "letters")

    // Grab the reply toys
    dataBytes, err := io.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // Show the reply
    content := string(dataBytes)
    fmt.Println("Reply content:\n", content)
}

// Helper for the home page, like a welcome sign
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to LearnCodeOnline Fruit Shop Server!")
}

// Helper for the /get page, like sending a message
func getHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "Hello from the fruit server!"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Helper for the /post page, like receiving a toy
func postHandler(w http.ResponseWriter, r *http.Request) {
    // Check for POST request
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed!", http.StatusMethodNotAllowed)
        return
    }

    // Read the toy box
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Can't read toy box!", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Try JSON data (user or course)
    var u user
    if err := json.Unmarshal(body, &u); err == nil && u.Name != "" {
        fmt.Printf("Got user JSON: Name=%s, Age=%d\n", u.Name, u.Age)
        response := map[string]string{"message": fmt.Sprintf("Received user %s, age %d!", u.Name, u.Age)}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
        return
    }

    var c course
    if err := json.Unmarshal(body, &c); err == nil && c.CourseName != "" {
        fmt.Printf("Got course JSON: Course=%s, Price=%d, Platform=%s\n", c.CourseName, c.Price, c.Platform)
        response := map[string]string{"message": fmt.Sprintf("Received course %s from %s!", c.CourseName, c.Platform)}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
        return
    }

    // Try form data
    if err := r.ParseForm(); err == nil {
        name := r.FormValue("name")
        ageStr := r.FormValue("age")
        if name != "" && ageStr != "" {
            fmt.Printf("Got form: Name=%s, Age=%s\n", name, ageStr)
            response := map[string]string{"message": fmt.Sprintf("Received form: %s, age %s!", name, ageStr)}
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(response)
            return
        }
    }

    // If all fail
    http.Error(w, "Need valid JSON (user/course) or form with name and age!", http.StatusBadRequest)
}

// Helper for the /postform page, like receiving a form note
func postFormHandler(w http.ResponseWriter, r *http.Request) {
    // Check for POST request
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed!", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Can't read form note!", http.StatusBadRequest)
        return
    }

    // Get form values
    fruit := r.FormValue("fruit")
    quantity := r.FormValue("quantity")
    customer := r.FormValue("customer")

    // Check if we got the right toys
    if fruit == "" || quantity == "" || customer == "" {
        http.Error(w, "Need fruit, quantity, and customer in form!", http.StatusBadRequest)
        return
    }

    // Show what we got
    fmt.Printf("Got form: Fruit=%s, Quantity=%s, Customer=%s\n", fruit, quantity, customer)
    response := map[string]string{
        "message": fmt.Sprintf("Received order: %s wants %s %s!", customer, quantity, fruit),
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Explanation comments for a 5-year-old:
// - What's a web server? It's like a fruit shop where people visit to see signs (pages) or send toys (data)!
// - What's a GET request? It's like sending a letter to the shop asking for a toy, like a message!
// - What's a POST request? It's like sending a toy box (JSON) or note (form) to the shop for them to keep!
// - What's a form POST? It's like sending a simple note with words (like "fruit=Mango") instead of a fancy toy box.
// - What's 'http.Get(myurl)'? It's like mailing a letter to an address (like http://localhost:8000/get) to get toys.
// - What's 'http.Post(myurl, "application/json", ...)'? It's like sending a toy box with JSON words (like a course order).
// - What's 'http.PostForm(myurl, formData)'? It's like sending a note with fruit order words (like "Mango, 3").
// - What's 'strings.NewReader(...)'? It's like making a toy box with JSON words, like a course order.
// - What's 'url.Values{...}'? It's like writing a note with fruit order toys, like "fruit=Mango, quantity=3".
// - What's 'defer response.Body.Close()'? It's like promising to shut the toy box after taking or sending toys.
// - What's 'response.StatusCode'? It's like a note saying if the shop is happy (200 is "OK") or not.
// - What's 'io.ReadAll(response.Body)'? It's like grabbing all the reply toys (words) from the shop.
// - What's 'string(dataBytes)'? It's like turning toys into a story we can read.
// - What's 'http.HandleFunc("/postform", postFormHandler)'? It's like setting up a special stall for form notes.
// - What's 'r.ParseForm()'? It's like reading a visitor's note to find words like "fruit" and "quantity".
// - What's 'r.FormValue("fruit")'? It's like looking for the "fruit" word on the note, like "Mango".
// - Why JSON and form? It's like letting visitors send fancy toy boxes (JSON) or simple notes (form).
// - Why user, course, and form? It lets the shop handle different toys: people (user), orders (course), or notes (form).
// - How to set up the server? 
//   1. Save this file in 21lcoWebServer/main.go.
//   2. Run 'go mod init 21lcoWebServer' in the folder.
//   3. Run 'go run main.go' to start the server at http://localhost:8000.
//   4. Keep it running, then test requests (below) or run the program again in another terminal.
// - How to test with ThunderClient? In VS Code, get ThunderClient:
//   1. GET: New Request > http://localhost:8000 or http://localhost:8000/get > Send. See welcome or {"message":"Hello from the fruit server!"}.
//   2. POST (User JSON): New Request > http://localhost:8000/post > POST > Body > JSON > {"name":"Hitesh","age":4} > Send. See {"message":"Received Hitesh, age 4!"}.
//   3. POST (Course JSON): New Request > http://localhost:8000/post > POST > Body > JSON > {"coursename":"Let's go with golang","price":0,"platform":"learncodeonline.in"} > Send. See {"message":"Received course Let's go with golang from learncodeonline.in!"}.
//   4. POST (Form): New Request > http://localhost:8000/post > POST > Body > Form > name=Hitesh, age=4 > Send. See {"message":"Received form: Hitesh, age 4!"}.
//   5. POST (Form to /postform): New Request > http://localhost:8000/postform > POST > Body > Form > fruit=Mango, quantity=3, customer=Justin > Send. See {"message":"Received order: Justin wants 3 Mango!"}.
// - Why avoid CodePen? CodePen is for web toys (HTML, JavaScript), not Go. Use VS Code with Go extension and ThunderClient; it runs 'go run main.go' and tests web shops easily.
// - Why fruit shop? It keeps the fun fruit theme from earlier chapters, like running a web shop.
// - How does this help beginners? It shows how to make a shop (server) and send/ask for toys (GET/POST requests) like a Node.js shop, with easy ThunderClient steps to test JSON and form notes.