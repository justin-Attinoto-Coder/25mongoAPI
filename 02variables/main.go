package main

import "fmt"

func main() {
    // Print chapter focus
    fmt.Println("Exploring Variables, Types, and Constants")

    // Variable declarations with explicit types
    var username string = "Justin"
    var age int = 41

    // Short variable declaration
    city := "New York"

    // Multiple variables
    var (
        height float32 = 5.9
        isStudent bool = false
    )

    // Constant declarations
    const birthYear int = 1984
    const loginToken string = "gibberish"

    // Short variable declaration for numberOfUser
    numberOfUser := 30000.0

    // Print variables and their types
    fmt.Println(username)
    fmt.Printf("variable is of type: %T \n", username)
    fmt.Printf("Age: %d (Type: %T)\n", age, age)
    fmt.Printf("City: %s (Type: %T)\n", city, city)
    fmt.Printf("Height: %.1f (Type: %T)\n", height, height)
    fmt.Printf("Is Student: %t (Type: %T)\n", isStudent, isStudent)
    fmt.Printf("Birth Year: %d (Type: %T)\n", birthYear, birthYear)
    fmt.Printf("Login Token: %s (Type: %T)\n", loginToken, loginToken)
    fmt.Println(numberOfUser)
    fmt.Printf("Number of Users: %.1f (Type: %T)\n", numberOfUser, numberOfUser)

    // Basic operation with variables
    yearsSinceBirth := 2025 - birthYear
    fmt.Printf("%s, you are %d years old in 2025.\n", username, yearsSinceBirth)
}