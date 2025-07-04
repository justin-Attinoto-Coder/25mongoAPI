package main

import (
    "fmt"
    "strings"
)

// A struct is like a custom toy box that holds different toys together
type user struct {
    name    string // The user's name, like "Justin"
    email   string // Their email, like "justin@example.com"
    status  bool   // True if email is verified, false if not
    age     int    // How old they are, like 41
    phone   string // Their phone number, like "555-1234"
    address string // Where they live, like "123 Fruit Street"
}

// The main function is like the start button for our program!
func main() {
    // Say hello to our struct adventure
    fmt.Println("Welcome to our custom toy box adventure with structs!")

    // Make a new user toy box for Justin
    justin := user{
        name:    "Justin",
        email:   "justin@example.com",
        status:  false,
        age:     41,
        phone:   "555-1234",
        address: "123 Fruit Street",
    }
    fmt.Println("Justin's toy box:", justin)

    // Use a helper to show Justin's details
    justin.showDetails()

    // Verify Justin's email with a helper
    justin.verifyEmail()
    fmt.Println("Justin's status after verification:", justin.status)

    // Update Justin's phone with a helper
    justin.updatePhone("555-9999")
    fmt.Println("Justin's new phone:", justin.phone)

    // Check how many years until Justin retires
    yearsToRetire := justin.yearsUntilRetirement(65)
    fmt.Println("Years until Justin retires:", yearsToRetire)

    // Make another user for Alice using a helper
    alice := createUser("Alice", "alice@example.com", 25, "555-5678", "456 Banana Avenue")
    fmt.Println("Alice's toy box:", alice)

    // Show Alice's details and verify her email
    alice.showDetails()
    alice.verifyEmail()
    fmt.Println("Alice's status after verification:", alice.status)

    // Make a list of users, like a shelf of toy boxes
    users := []user{justin, alice}
    fmt.Println("All user toy boxes:", users)

    // Loop through the shelf and use helpers for each user
    fmt.Println("Checking all users with helpers:")
    for i, u := range users {
        fmt.Printf("User %d:\n", i)
        u.showDetails()
        fmt.Println("Years to retire:", u.yearsUntilRetirement(65))
    }
}

// This helper makes a new user toy box
func createUser(name, email string, age int, phone, address string) user {
    return user{
        name:    name,
        email:   email,
        status:  false,
        age:     age,
        phone:   phone,
        address: address,
    }
}

// This helper shows a user's details, like opening the toy box
func (u user) showDetails() {
    fmt.Printf("Name: %s, Email: %s, Verified: %t, Age: %d, Phone: %s, Address: %s\n",
        u.name, u.email, u.status, u.age, u.phone, u.address)
}

// This helper verifies a user's email, like stamping the box
func (u *user) verifyEmail() {
    if strings.Contains(u.email, "@") && strings.HasSuffix(u.email, ".com") {
        u.status = true // Stamp as verified
    }
}

// This helper changes a user's phone, like swapping a toy
func (u *user) updatePhone(newPhone string) {
    u.phone = newPhone // Put the new phone in the box
}

// This helper counts years until retirement, like a future calculator
func (u user) yearsUntilRetirement(retireAge int) int {
    years := retireAge - u.age
    if years < 0 {
        return 0 // If already past retire age, say 0
    }
    return years
}

// Explanation comments for a 5-year-old:
// - What's a struct? It's like a custom toy box where you put different toys (like name, email, age) together for one person.
// - What's 'type user struct'? It's like naming our toy box "user" and saying what toys it can hold (name, email, status, etc.).
// - What's a method? It's like a special helper that knows how to play with the toy box, like showing toys or changing them.
// - What's 'justin := user{...}'? It's like making a new toy box for Justin with his toys (name="Justin", age=41).
// - What's 'justin.showDetails()'? It's like asking a helper to open Justin's box and show all his toys (name, email, etc.).
// - What's 'justin.verifyEmail()'? It's like a helper checking Justin's email and stamping his box "verified" (true).
// - What's 'justin.updatePhone("555-9999")'? It's like a helper swapping Justin's phone toy for a new one.
// - What's 'justin.yearsUntilRetirement(65)'? It's like a helper counting how many years until Justin is 65 (retirement age).
// - What's 'func (u user) showDetails()'? It's a helper that looks at a user box without changing it (uses 'u' not '*u').
// - What's 'func (u *user) verifyEmail()'? It's a helper that can change the box (uses '*u' to swap toys like status).
// - What's 'createUser(...)'? It's like a toy box maker that builds a new user box with the toys you give it.
// - What's 'users := []user{justin, alice}'? It's like putting Justin's and Alice's toy boxes on a shelf.
// - What's 'for i, u := range users'? It's like a magic hand picking up each toy box and showing its toys.
// - Why no inheritance? Go doesn't let toy boxes copy from other boxes (no "parent" or "super"). Each box is unique, but helpers (methods) make it special.
// - Why phone and address? These extra toys make the user box more fun and show structs can hold lots of things, like in real apps.
// - How does this help beginners? It shows how to make toy boxes (structs) and use helpers (methods) with simple examples like email verification, keeping the fruit theme (like "Fruit Street") for fun.