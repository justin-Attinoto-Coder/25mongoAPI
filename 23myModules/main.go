package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our module adventure
    fmt.Println("Welcome to our Go module adventure!")
    greeter()

    // Set up a fruit shop web server with routes
    router := mux.NewRouter()
    router.HandleFunc("/", serveHome)               // Home page
    router.HandleFunc("/fruits", listFruits)        // List fruits
    router.HandleFunc("/fruit/{name}", getFruit)    // Get one fruit

    // Start the shop at port 8000
    fmt.Println("Starting fruit shop server at http://localhost:8000")
    err := http.ListenAndServe(":8000", router)
    if err != nil {
        panic(err)
    }
}

// Helper to say hello, like waving to friends
func greeter() {
    fmt.Println("Hello mod in golang!")
}

// Helper for the home page, like a welcome sign
func serveHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

// Helper to list fruits, like showing a fruit menu
func listFruits(w http.ResponseWriter, r *http.Request) {
    fruits := []string{"Apple", "Banana", "Mango"}
    w.Write([]byte("<h1>Our Fruits</h1><ul>"))
    for _, fruit := range fruits {
        fmt.Fprintf(w, "<li>%s</li>", fruit)
    }
    w.Write([]byte("</ul>"))
}

// Helper to get one fruit, like picking a fruit by name
func getFruit(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fruitName := vars["name"]
    fmt.Fprintf(w, "<h1>You picked %s!</h1>", fruitName)
}

// Explanation comments for a 5-year-old:
// - What's a Go module? It's like a big toy box with a name and number that holds all your toys (code) and their friends (dependencies).
// - What's 'module github.com/hiteshchoudhary/myModules'? It's like naming your toy box "myModules" so everyone knows where to find it.
// - What's 'go 1.17'? It's like saying your toy box works with Go version 1.17 or newer.
// - What's semantic versioning? It's like giving your toy box a number like v1.2.3: v1 (big changes), 2 (new toys), 3 (fixes). See semver.org for more![](https://medium.com/%40emonemrulhasan35/go-modules-0a6ce56b4209)
// - What's 'go mod init github.com/hiteshchoudhary/myModules'? It's like making a new toy box with a name and a go.mod file to list toys.
// - What's 'go mod tidy'? It's like cleaning your toy box to keep only the toys you need and throw out extras.[](https://www.pingcap.com/article/understanding-go-modules-for-beginners/)
// - What's 'go mod why github.com/gorilla/mux'? It's like asking, "Why do we need this gorilla/mux toy?" to see who uses it.[](https://blog.devtrovert.com/p/go-get-go-mod-tidy-commands)
// - What's 'github.com/gorilla/mux'? It's like a super helper toy for making web shop stalls (routes) to show pages.[](https://pkg.go.dev/github.com/gorilla/mux)
// - What's 'go get github.com/gorilla/mux'? It's like grabbing the gorilla/mux toy from the internet and adding it to your toy box.
// - What's 'func greeter()'? It's like a helper that waves and says "Hello mod in golang!" to start the fun.
// - What's 'func serveHome(w, r)'? It's like putting a welcome sign at the shop's front door for visitors.
// - What's 'router := mux.NewRouter()'? It's like building a shop with stalls where visitors can go (routes).
// - What's 'router.HandleFunc("/fruits", listFruits)'? It's like making a stall to show a fruit menu.
// - What's 'router.HandleFunc("/fruit/{name}", getFruit)'? It's like a stall where you pick a fruit by name, like "Mango".
// - What's 'mux.Vars(r)'? It's like checking the visitor's note to see which fruit they want.
// - How to set up? 
//   1. Make folder '23myModules'.
//   2. Run 'go mod init github.com/hiteshchoudhary/myModules' to create go.mod.
//   3. Run 'go get github.com/gorilla/mux' to add the router toy.
//   4. Save this main.go file.
//   5. Run 'go mod tidy' to clean the toy box.
//   6. Run 'go mod why github.com/gorilla/mux' to see why we need it.
//   7. Run 'go run main.go' to start the server.
// - How to test? Visit http://localhost:8000, http://localhost:8000/fruits, or http://localhost:8000/fruit/Mango in a browser. Use ThunderClient in VS Code for more tests.
// - Why fruit shop? It keeps the fun fruit theme (like Apple, Mango) from earlier chapters, like running a web shop.
// - How does this help beginners? It shows how to make a toy box (module), add helper toys (gorilla/mux), and build a web shop with routes, like organizing and sharing toys in Go.