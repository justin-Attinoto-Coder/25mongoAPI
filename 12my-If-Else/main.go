package main

import (
    "fmt"
    "strings"
)

// A struct is like a custom toy box for user info
type user struct {
    name   string // User's name, like "Justin"
    email  string // Their email, like "justin@example.com"
    status bool   // True if email is verified
    age    int    // How old they are, like 41
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our if-else adventure
    fmt.Println("Welcome to our decision-making adventure with if and else!")

    // Make a user toy box for Justin
    justin := user{
        name:   "Justin",
        email:  "justin@example.com",
        status: false,
        age:    41,
    }

    // Check if Justin is old enough to vote (18 or older)
    if justin.age >= 18 {
        fmt.Println("Justin can vote! He's", justin.age, "years old.")
    } else {
        fmt.Println("Justin is too young to vote. He's only", justin.age)
    }

    // Check if Justin's email is verified
    if justin.status {
        fmt.Println("Justin's email is verified! Welcome!")
    } else {
        fmt.Println("Justin, please verify your email!")
        justin.verifyEmail() // Try verifying it
        fmt.Println("After verification, status is:", justin.status)
    }

    // Check Justin's age group with multiple choices
    if justin.age < 13 {
        fmt.Println("Justin is a kid!")
    } else if justin.age < 20 {
        fmt.Println("Justin is a teen!")
    } else if justin.age < 30 {
        fmt.Println("Justin is a young adult!")
    } else {
        fmt.Println("Justin is an adult!")
    }

    // Make another user, Alice
    alice := user{
        name:   "Alice",
        email:  "alice@example.com",
        status: true,
        age:    25,
    }

    // Check if Alice has a fruit-themed email
    if strings.Contains(alice.email, "fruit") {
        fmt.Println("Alice has a fruity email!")
    } else {
        fmt.Println("Alice's email isn't fruity!")
    }

    // Track login attempts, like counting wrong password tries
    loginAttempts := 6 // Pretend Justin tried 6 times
    maxAttempts := 5   // Only allow 5 tries
    if loginAttempts <= maxAttempts {
        fmt.Println("Login okay! You tried", loginAttempts, "times.")
    } else {
        fmt.Println("Too many tries!", loginAttempts, "is more than", maxAttempts)
        fmt.Println("Account locked! Call Fruit Support at 555-FRUIT!")
    }

    // Combine conditions: Check if Alice is verified AND young
    if alice.status && alice.age < 30 {
        fmt.Println("Alice is verified and young! Special fruit discount!")
    } else {
        fmt.Println("Alice doesn't get the young verified discount.")
    }

    // Check users in a list, like a shelf of toy boxes
    users := []user{justin, alice}
    for i, u := range users {
        fmt.Printf("Checking user %d:\n", i)
        if u.status {
            fmt.Println(u.name, "is verified!")
        } else {
            fmt.Println(u.name, "needs to verify their email!")
        }
    }
}

// This helper verifies a user's email, like stamping the box
func (u *user) verifyEmail() {
    if strings.Contains(u.email, "@") && strings.HasSuffix(u.email, ".com") {
        u.status = true // Stamp as verified
    }
}

// Explanation comments for a 5-year-old:
// - What's an if statement? It's like a decision helper that checks if something is true, like "Is Justin old enough?" If yes, it does one thing!
// - What's an else statement? It's like saying, "If not, do this instead!" Like if Justin isn't old enough, say he's too young.
// - What's 'if justin.age >= 18'? It's like asking, "Is Justin's age 18 or more?" If yes, he can vote!
// - What's 'else if justin.age < 20'? It's like checking another question if the first one isn't true, like "Is Justin a teen?"
// - What's 'loginAttempts <= maxAttempts'? It's like counting how many times someone tried a password and checking if it's too many (more than 5).
// - What's 'if alice.status && alice.age < 30'? It's like asking two questions together: "Is Alice verified AND young?" Both must be true!
// - What's 'justin.verifyEmail()'? It's a helper that checks Justin's email and stamps his toy box as verified.
// - Why check login attempts? It's like a watchout to stop someone from guessing passwords too many times (more than 5), locking their account to keep it safe.
// - Why use a user struct? It's like a toy box to keep user toys (name, email, age) together, making it easy to check them with if-else.
// - Why fruit theme? It connects to earlier chapters (like "Fruit Street") to make it fun, and we used "Fruit Support" for the lockout message.
// - How does this help beginners? It shows how to make decisions with if, else, and else if, using simple checks like age, email, and login tries, plus a cool lockout feature to learn safety.