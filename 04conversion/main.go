package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Initialize module (run in terminal: go mod init conversion)

    // Print welcome message
    fmt.Println("Welcome to our pizza app")
    fmt.Println("Please rate our pizza between 1 and 5")

    // Create reader for user input
    reader := bufio.NewReader(os.Stdin)

    // Read pizza rating
    fmt.Print("Enter rating: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    fmt.Println("Thanks for rating, ", input)

    // Convert string to float64
    numRating, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Error converting rating to float:", err)
        panic(err)
    }
    fmt.Printf("Added 1 to your rating: %.1f\n", numRating+1)

    // Convert float to string for output
    ratingStr := strconv.FormatFloat(numRating, 'f', 1, 64)
    fmt.Println("Your rating as string:", ratingStr)

    // Read number of pizzas ordered (demonstrate string to int conversion)
    fmt.Print("Enter number of pizzas ordered: ")
    pizzaInput, _ := reader.ReadString('\n')
    pizzaInput = strings.TrimSpace(pizzaInput)

    numPizzas, err := strconv.Atoi(pizzaInput) // Atoi for string to int
    if err != nil {
        fmt.Println("Error converting to integer:", err)
        panic(err)
    }
    fmt.Printf("You ordered %d pizzas\n", numPizzas)

    // Convert int to string
    pizzaCountStr := strconv.Itoa(numPizzas)
    fmt.Println("Number of pizzas as string:", pizzaCountStr)

    // Read if customer is premium (demonstrate string to bool conversion)
    fmt.Print("Are you a premium member? (true/false): ")
    premiumInput, _ := reader.ReadString('\n')
    premiumInput = strings.TrimSpace(premiumInput)

    isPremium, err := strconv.ParseBool(premiumInput)
    if err != nil {
        fmt.Println("Error converting to boolean:", err)
        panic(err)
    }
    fmt.Printf("Premium member status: %t\n", isPremium)

    // Convert bool to string
    premiumStr := strconv.FormatBool(isPremium)
    fmt.Println("Premium status as string:", premiumStr)

    // Demonstrate type assertion with comma ok syntax
    var data interface{} = numRating
    if val, ok := data.(float64); ok {
        fmt.Printf("Rating is a float64: %.1f\n", val)
    } else {
        fmt.Println("Rating is not a float64")
    }
}