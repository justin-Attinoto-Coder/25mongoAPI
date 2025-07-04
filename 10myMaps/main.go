package main

import "fmt"

// The main function is like the start button for our program!
func main() {
    // Say hello to our map adventure
    fmt.Println("Welcome to our magic toy box adventure with maps!")

    // A map is like a magic toy box where each toy has a special name tag
    // Let's make a box to keep track of how many fruits we have
    fruitInventory := make(map[string]int) // A box for fruit names and their counts
    // "string" is the name tag (key), "int" is the toy (value, like number of fruits)

    // Put some fruits in the magic box with their counts
    fruitInventory["Apple"] = 10  // Tag "Apple" gets 10 apples
    fruitInventory["Banana"] = 15 // Tag "Banana" gets 15 bananas
    fruitInventory["Mango"] = 8   // Tag "Mango" gets 8 mangos
    fmt.Println("My fruit inventory:", fruitInventory)
    // This shows: map[Apple:10 Banana:15 Mango:8]

    // Loop through the fruit box to see all toys and their tags
    fmt.Println("Checking all fruits with a magic hand:")
    for fruit, count := range fruitInventory {
        fmt.Println("Tag", fruit, "has", count, "fruits")
    }
    // This shows:
    // Tag Apple has 10 fruits
    // Tag Banana has 15 fruits
    // Tag Mango has 8 fruits

    // Check how many bananas we have by looking at the "Banana" tag
    bananaCount := fruitInventory["Banana"]
    fmt.Println("Number of bananas:", bananaCount)
    // This shows: 15

    // Check if a fruit is in the box using a magic trick (comma ok)
    count, exists := fruitInventory["Mango"] // Ask if "Mango" is there
    if exists {
        fmt.Println("Found Mango with", count, "items!")
    } else {
        fmt.Println("No Mango in the box!")
    }
    // This shows: Found Mango with 8 items!

    // Remove a fruit from the box, like taking a toy out
    delete(fruitInventory, "Mango") // Remove the "Mango" tag and its toy
    fmt.Println("After removing Mango:", fruitInventory)
    // This shows: map[Apple:10 Banana:15]

    // Make a box for student scores, like grades in a class
    studentScores := map[string]float64{ // A box with student names and scores
        "Justin": 85.5, // Justin gets 85.5
        "Alice":  92.0, // Alice gets 92.0
        "Bob":    78.5, // Bob gets 78.5
    }
    fmt.Println("Student scores:", studentScores)
    // This shows: map[Alice:92 Bob:78.5 Justin:85.5]

    // Loop through student scores to see all grades
    fmt.Println("Checking all student scores with a magic hand:")
    for student, score := range studentScores {
        fmt.Println("Student", student, "has score", score)
    }
    // This shows:
    // Student Justin has score 85.5
    // Student Alice has score 92
    // Student Bob has score 78.5

    // Add a new student score
    studentScores["Emma"] = 88.0 // Add Emma with 88.0
    fmt.Println("After adding Emma:", studentScores)
    // This shows: map[Alice:92 Bob:78.5 Emma:88 Justin:85.5]

    // Make a box for a course schedule, like class times
    courseSchedule := make(map[string]string)
    courseSchedule["ReactJS"] = "Monday 9AM"    // ReactJS class on Monday
    courseSchedule["Python"] = "Tuesday 2PM"    // Python class on Tuesday
    courseSchedule["Swift"] = "Wednesday 11AM"  // Swift class on Wednesday
    fmt.Println("Course schedule:", courseSchedule)
    // This shows: map[Python:Tuesday 2PM ReactJS:Monday 9AM Swift:Wednesday 11AM]

    // Loop through course schedule to see all class times
    fmt.Println("Checking all courses with a magic hand:")
    for course, time := range courseSchedule {
        fmt.Println("Course", course, "is at", time)
    }
    // This shows:
    // Course ReactJS is at Monday 9AM
    // Course Python is at Tuesday 2PM
    // Course Swift is at Wednesday 11AM

    // Change a class time
    courseSchedule["Swift"] = "Friday 10AM" // Move Swift to Friday
    fmt.Println("Updated course schedule:", courseSchedule)
    // This shows: map[Python:Tuesday 2PM ReactJS:Monday 9AM Swift:Friday 10AM]

    // Make a new box for programming languages, like a list of coding toys
    languages := map[string]string{ // A box with language names and their types
        "Go":         "Compiled",    // Go is a compiled language
        "Python":     "Interpreted", // Python is interpreted
        "JavaScript": "Interpreted", // JavaScript is interpreted
        "Ruby":       "Interpreted", // Ruby is interpreted
        "Swift":      "Compiled",    // Swift is compiled
    }
    fmt.Println("Programming languages:", languages)
    // This shows: map[Go:Compiled JavaScript:Interpreted Python:Interpreted Ruby:Interpreted Swift:Compiled]

    // Loop through languages to see each one, like checking coding toys
    fmt.Println("Checking all languages with a magic hand:")
    for lang, langType := range languages {
        fmt.Println("Language", lang, "is", langType)
    }
    // This shows:
    // Language Go is Compiled
    // Language Python is Interpreted
    // Language JavaScript is Interpreted
    // Language Ruby is Interpreted
    // Language Swift is Compiled

    // Add a new language to the box
    languages["Java"] = "Compiled" // Add Java as a compiled language
    fmt.Println("After adding Java:", languages)
    // This shows: map[Go:Compiled Java:Compiled JavaScript:Interpreted Python:Interpreted Ruby:Interpreted Swift:Compiled]

    // Check if a language is in the box
    langType, exists := languages["Python"] // Ask if "Python" is there
    if exists {
        fmt.Println("Found Python, it's", langType, "!")
    } else {
        fmt.Println("No Python in the box!")
    }
    // This shows: Found Python, it's Interpreted !

    // Remove a language from the box
    delete(languages, "Ruby") // Remove Ruby
    fmt.Println("After removing Ruby:", languages)
    // This shows: map[Go:Compiled Java:Compiled JavaScript:Interpreted Python:Interpreted Swift:Compiled]

    // Count how many languages are in the box
    langCount := len(languages)
    fmt.Println("Number of languages in the box:", langCount)
    // This shows: 5

    // Memory tip: Maps are magic boxes that grow as needed
    // But keep them tidy by only adding what you need!
    fmt.Println("All our magic boxes are ready to play!")
}