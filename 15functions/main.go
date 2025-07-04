package main

import "fmt"

// The main function is like the start button for our program!
func main() {
    // Say hello to our function adventure
    fmt.Println("Welcome to our fruit shop function adventure!")

    // Call a simple helper to say hi
    sayHello()
    // This shows: Hello from the fruit shop!

    // Greet a user with their name
    greetUser("Justin")
    // This shows: Welcome, Justin, to the fruit shop!

    // Calculate the price of apples
    applePrice := getFruitPrice("Apple", 2)
    fmt.Println("Price for 2 apples:", applePrice, "coins")
    // This shows: Price for 2 apples: 4 coins

    // Check if we have enough bananas
    haveBananas := checkStock("Banana", 5)
    fmt.Println("Enough bananas for 5?", haveBananas)
    // This shows: Enough bananas for 5? true

    // Get total and discount for a fruit order
    total, discount := calculateOrder("Mango", 10)
    fmt.Println("Total for 10 mangos:", total, "coins, Discount:", discount, "coins")
    // This shows: Total for 10 mangos: 30 coins, Discount: 3 coins

    // Buy lots of fruits with a flexible helper
    totalCost := buyFruits("Apple", "Banana", "Mango")
    fmt.Println("Total cost for fruits:", totalCost, "coins")
    // This shows: Total cost for fruits: 12 coins

    // Use a quick helper to count fruits
    countFruits := func() int {
        return 3 // Pretend we have 3 fruits
    }
    fmt.Println("Number of fruit types:", countFruits())
    // This shows: Number of fruit types: 3
}

// This helper says hello, like waving at everyone
func sayHello() {
    fmt.Println("Hello from the fruit shop!")
}

// This helper greets a user, like saying hi to a friend
func greetUser(name string) {
    fmt.Println("Welcome,", name, ", to the fruit shop!")
}

// This helper gets the price for a fruit, like checking a price tag
func getFruitPrice(fruit string, quantity int) int {
    pricePerUnit := 2 // Each fruit costs 2 coins
    return pricePerUnit * quantity
}

// This helper checks if we have enough fruit, like looking in the shop
func checkStock(fruit string, needed int) bool {
    stock := map[string]int{"Apple": 10, "Banana": 8, "Mango": 5}
    return stock[fruit] >= needed
}

// This helper calculates total and discount, like adding up a bill
func calculateOrder(fruit string, quantity int) (int, int) {
    price := getFruitPrice(fruit, quantity) // Use another helper
    discount := 0
    if quantity > 5 {
        discount = price / 10 // 10% off for big orders
    }
    return price, discount
}

// This helper adds up costs for many fruits, like a big shopping list
func buyFruits(fruits ...string) int {
    total := 0
    for _, fruit := range fruits {
        total += getFruitPrice(fruit, 2) // Buy 2 of each fruit
    }
    return total
}

// Explanation comments for a 5-year-old:
// - What's a function? It's like a special helper that does a job, like saying hi or counting fruit coins!
// - What's 'func sayHello()'? It's a helper that waves and says hello without needing any toys.
// - What's 'func greetUser(name string)'? It's like a helper that says hi to a friend using their name toy.
// - What's 'func getFruitPrice(fruit string, quantity int) int'? It's like a helper that checks a fruit's price tag and gives back a number (coins).
// - What's 'func checkStock(fruit string, needed int) bool'? It's like a helper that looks in the shop to see if we have enough fruit toys, saying yes (true) or no (false).
// - What's 'func calculateOrder(fruit string, quantity int) (int, int)'? It's like a helper that adds up a bill and gives back two toys: total coins and a discount.
// - What's 'func buyFruits(fruits ...string) int'? It's like a helper that takes any number of fruit toys and adds up their prices, like a big shopping list.
// - What's 'countFruits := func() int'? It's like a quick helper we make on the spot to count fruit types, without giving it a big name.
// - Why use functions? They make jobs easy, like having helpers for each task in the fruit shop, so we don't mix up everything.
// - Why fruit shop? It keeps the fruit theme (like Apple, Mango) from earlier chapters, making it fun to learn with a shop game.
// - How does this help beginners? It shows all kinds of helpers (functions) with simple jobs like greeting, pricing, and checking stock, so you learn how to make and use them in a fruit shop adventure.