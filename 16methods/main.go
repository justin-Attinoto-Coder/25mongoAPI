package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

// A user toy box holds info about a person
type user struct {
    name    string // User's name, like "Justin"
    email   string // Their email, like "justin@example.com"
    status  bool   // True if email is verified
    age     int    // How old they are, like 41
}

// A fruitBasket toy box holds fruits and their counts
type fruitBasket struct {
    fruits map[string]int // A magic box for fruit names and counts
    name   string         // Name of the basket, like "Justin's Basket"
}

// A gamePlayer toy box tracks a player's dice game
type gamePlayer struct {
    name  string // Player's name, like "Alice"
    score int    // Their score, like 10
    rolls int    // Number of dice rolls, like 3
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our method adventure
    fmt.Println("Welcome to our toy box helper adventure with methods!")

    // Example 1: User profile methods
    fmt.Println("\nExample 1: User Profile")
    justin := user{
        name:   "Justin",
        email:  "justin@example.com",
        status: false,
        age:    41,
    }
    justin.showProfile()        // Show Justin's info
    justin.verifyEmail()        // Verify his email
    justin.updateAge(42)        // Update his age
    fmt.Println("After updates:")
    justin.showProfile()

    // Example 2: Fruit basket methods
    fmt.Println("\nExample 2: Fruit Basket")
    basket := fruitBasket{
        fruits: map[string]int{"Apple": 5, "Banana": 3},
        name:   "Justin's Basket",
    }
    basket.showContents()         // Show fruits
    basket.addFruit("Mango", 4)   // Add mangos
    basket.removeFruit("Banana")  // Remove bananas
    fmt.Println("After updates:")
    basket.showContents()

    // Example 3: Dice game player methods
    fmt.Println("\nExample 3: Dice Game Player")
    rand.Seed(time.Now().UnixNano()) // Mix dice for random rolls
    alice := gamePlayer{
        name:  "Alice",
        score: 0,
        rolls: 0,
    }
    alice.showStatus()           // Show Alice's score
    alice.rollDice()             // Roll a dice
    alice.doubleScore()          // Double her score
    fmt.Println("After playing:")
    alice.showStatus()
}

// Helper to show a user's profile
func (u user) showProfile() {
    fmt.Printf("Profile: Name=%s, Email=%s, Verified=%t, Age=%d\n",
        u.name, u.email, u.status, u.age)
}

// Helper to verify a user's email
func (u *user) verifyEmail() {
    if strings.Contains(u.email, "@") && strings.HasSuffix(u.email, ".com") {
        u.status = true
        fmt.Println(u.name, "email verified!")
    } else {
        fmt.Println(u.name, "email not valid!")
    }
}

// Helper to update a user's age
func (u *user) updateAge(newAge int) {
    u.age = newAge
    fmt.Println(u.name, "age updated to", newAge)
}

// Helper to show a fruit basket's contents
func (b fruitBasket) showContents() {
    fmt.Println(b.name, "contains:")
    for fruit, count := range b.fruits {
        fmt.Println(fruit, ":", count)
    }
}

// Helper to add fruit to a basket
func (b *fruitBasket) addFruit(fruit string, count int) {
    b.fruits[fruit] += count
    fmt.Println("Added", count, fruit, "to", b.name)
}

// Helper to remove fruit from a basket
func (b *fruitBasket) removeFruit(fruit string) {
    delete(b.fruits, fruit)
    fmt.Println("Removed", fruit, "from", b.name)
}

// Helper to show a game player's status
func (p gamePlayer) showStatus() {
    fmt.Printf("Player %s: Score=%d, Rolls=%d\n", p.name, p.score, p.rolls)
}

// Helper to roll a dice and add to score
func (p *gamePlayer) rollDice() {
    roll := rand.Intn(6) + 1 // Roll 1 to 6
    p.score += roll
    p.rolls++
    fmt.Println(p.name, "rolled a", roll, "!")
}

// Helper to double a player's score
func (p *gamePlayer) doubleScore() {
    p.score *= 2
    fmt.Println(p.name, "score doubled to", p.score)
}

// Explanation comments for a 5-year-old:
// - What's a method? It's like a special helper that knows how to play with a toy box (like a user or fruit basket) and do jobs for it.
// - What's 'type user struct'? It's like a toy box named "user" that holds toys like name, email, and age.
// - What's 'func (u user) showProfile()'? It's a helper that opens the user box and shows all its toys without changing them.
// - What's 'func (u *user) verifyEmail()'? It's a helper that can change the user box (like stamping it verified) using a magic map (*u).
// - What's 'justin.showProfile()'? It's like asking the helper to show Justin's toys (name, email, etc.).
// - What's 'type fruitBasket struct'? It's a toy box for fruits, with a magic box inside to hold fruit names and counts.
// - What's 'basket.addFruit("Mango", 4)'? It's like telling the helper to put 4 mango toys in the basket.
// - What's 'type gamePlayer struct'? It's a toy box for a game player, holding their name, score, and dice rolls.
// - What's 'alice.rollDice()'? It's like asking the helper to roll a dice for Alice and add the number to her score.
// - Why three examples? They show different jobs: user helpers (like checking emails), fruit helpers (like adding mangos), and game helpers (like rolling dice).
// - Why fruit theme? It keeps the fun from earlier chapters (like Apple, Mango) so the adventure feels connected.
// - How does this help beginners? It shows how helpers (methods) make toy boxes (structs) do cool things, like updating profiles, managing fruits, or playing games, with simple, fun examples.