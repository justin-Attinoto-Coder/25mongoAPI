package main

import (
    "fmt"
    "net/url"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our URL adventure
    fmt.Println("Welcome to handling URLs in golang!")

    // Example 1: Parse the provided URL
    const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
    fmt.Println("\nParsing URL:", myurl)
    parseURL(myurl)

    // Example 2: Parse a fruit shop URL with query params
    const fruitUrl = "https://fruitshop.com:8080/order?fruit=apple&quantity=5&customer=Justin"
    fmt.Println("\nParsing fruit shop URL:", fruitUrl)
    parseURL(fruitUrl)

    // Example 3: Build a new fruit shop URL
    fmt.Println("\nBuilding a new fruit shop URL")
    newUrl := buildFruitShopURL("https", "fruitshop.com", "/cart", map[string]string{
        "fruit":     "mango",
        "quantity":  "3",
        "promo":     "FRUIT10",
    })
    fmt.Println("New URL:", newUrl)
}

// Helper to parse a URL, like reading a website address
func parseURL(urlStr string) {
    // Break the address into parts
    result, err := url.Parse(urlStr)
    if err != nil {
        panic(err)
    }

    // Show the parts of the address
    fmt.Println("Scheme (like https):", result.Scheme)
    fmt.Println("Host (like lco.dev):", result.Host)
    fmt.Println("Path (like /learn):", result.Path)
    fmt.Println("Port (like 3000):", result.Port())
    fmt.Println("Raw Query (like coursename=reactjs):", result.RawQuery)

    // Get the query parts, like toys in a box
    qparams := result.Query()
    fmt.Printf("Query params are of type: %T\n", qparams)

    // Show specific query toys
    fmt.Println("Course/Fruit param:", qparams.Get("coursename"), qparams.Get("fruit"))

    // Loop through all query toys
    fmt.Println("All query params:")
    for key, values := range qparams {
        for _, val := range values {
            fmt.Printf("Param %s: %s\n", key, val)
        }
    }
}

// Helper to build a new URL, like making a website address
func buildFruitShopURL(scheme, host, path string, queryParams map[string]string) string {
    // Make a new address toy box
    partsOfUrl := &url.URL{
        Scheme: scheme,
        Host:   host,
        Path:   path,
    }

    // Add query toys to the address
    q := url.Values{}
    for key, value := range queryParams {
        q.Add(key, value)
    }
    partsOfUrl.RawQuery = q.Encode()

    // Turn the address into words
    return partsOfUrl.String()
}

// Explanation comments for a 5-year-old:
// - What's a URL? It's like a website address, like a map to a toy shop (fruitshop.com)!
// - What's 'url.Parse(myurl)'? It's like opening the address map to see its parts (like https or /learn).
// - What's 'result.Scheme'? It's like checking if the map says "https" or "http" for the shop.
// - What's 'result.Host'? It's like seeing the shop's name on the map, like "lco.dev".
// - What's 'result.Path'? It's like finding the shop section on the map, like "/learn".
// - What's 'result.Port()'? It's like a special door number for the shop, like "3000".
// - What's 'result.RawQuery'? It's like extra notes on the map, like "coursename=reactjs".
// - What's 'qparams := result.Query()'? It's like taking the extra notes and putting them in a toy box.
// - What's 'qparams.Get("coursename")'? It's like looking in the toy box for a specific note, like "reactjs".
// - What's 'for key, values := range qparams'? It's like a magic hand checking each note in the toy box.
// - What's 'partsOfUrl := &url.URL{...}'? It's like building a new map for a shop with parts like "https" and "fruitshop.com".
// - What's 'q.Add(key, value)'? It's like adding extra notes (like "fruit=mango") to the new map.
// - What's 'partsOfUrl.String()'? It's like turning the new map into words, like "https://fruitshop.com/cart?fruit=mango".
// - Why fruit shop? It keeps the fun fruit theme (like apple, mango) from earlier chapters, like visiting a web shop.
// - How does this help beginners? It shows how to read and make website maps (URLs) with simple examples, like finding course names or building fruit orders, making URLs easy to learn.