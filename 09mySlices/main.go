package main

import (
    "fmt"
    "math/rand"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our slice adventure
    fmt.Println("Welcome to our slice adventure with fruits, scores, and courses!")

    // A slice is like a stretchy toy box that can grow or shrink
    // Let's make a stretchy box for fruits
    fruitBasket := []string{"Apple", "Tomato", "Peach"} // Starts with 3 fruits
    fmt.Println("My stretchy fruit basket:", fruitBasket)
    // This shows: [Apple Tomato Peach]

    // Add more fruits to the stretchy box, like putting in new toys
    fruitBasket = append(fruitBasket, "Mango", "Banana") // Add Mango and Banana
    fmt.Println("After adding Mango and Banana:", fruitBasket)
    // This shows: [Apple Tomato Peach Mango Banana]

    // Remove a fruit (Tomato) by skipping it, like taking a toy out
    fruitBasket = append(fruitBasket[:1], fruitBasket[2:]...) // Keep before and after slot 1
    fmt.Println("Fruit basket after removing Tomato:", fruitBasket)
    // This shows: [Apple Peach Mango Banana]

    // Make a stretchy box for high scores, like points in a game
    highScores := []int{ // Start with 4 random numbers between 1 and 1000
        rand.Intn(1000) + 1, // Pick a random number, like rolling a dice
        rand.Intn(1000) + 1, // Another random number
        rand.Intn(1000) + 1, // Another one
        rand.Intn(1000) + 1, // One more
    }
    fmt.Println("My game high scores:", highScores)
    // This shows 4 random numbers, like: [542 317 896 123]

    // Add a new high score, like getting a new game score
    newScore := rand.Intn(1000) + 1 // Pick another random score
    highScores = append(highScores, newScore) // Add it to the box
    fmt.Println("After adding new score", newScore, ":", highScores)
    // This shows: [542 317 896 123 789]

    // Remove the second score (index 1), like taking out a score
    highScores = append(highScores[:1], highScores[2:]...) // Skip slot 1
    fmt.Println("High scores after removing second score:", highScores)
    // This shows: [542 896 123 789]

    // Make a stretchy box for courses, like a list of classes
    courses := []string{"ReactJS", "JavaScript", "Swift", "Python", "Ruby"}
    fmt.Println("My courses box:", courses)
    // This shows: [ReactJS JavaScript Swift Python Ruby]

    // Add a new course with "index +1" idea, like naming a course with a number
    // Let's say we add a course called "Course1" based on index
    newCourseIndex := len(courses) // Current number of courses (5)
    newCourse := fmt.Sprintf("Course%d", newCourseIndex+1) // Make "Course6"
    courses = append(courses, newCourse) // Add the new course
    fmt.Println("After adding", newCourse, ":", courses)
    // This shows: [ReactJS JavaScript Swift Python Ruby Course6]

    // Remove a course (Swift, index 2), like dropping a class
    courses = append(courses[:2], courses[3:]...) // Skip slot 2
    fmt.Println("Courses after removing Swift:", courses)
    // This shows: [ReactJS JavaScript Python Ruby Course6]

    // Check size and room in the courses box
    coursesSize := len(courses)
    coursesCapacity := cap(courses)
    fmt.Println("Courses box has", coursesSize, "courses and room for", coursesCapacity)
    // This shows: 5 courses, maybe room for 8 (capacity grows)

    // Look at one course by checking a slot
    fmt.Println("Course in slot 0:", courses[0])
    // This shows: ReactJS

    // Change a course, like renaming a class
    courses[1] = "NodeJS" // Change JavaScript to NodeJS
    fmt.Println("New courses box:", courses)
    // This shows: [ReactJS NodeJS Python Ruby Course6]

    // Take a piece of the courses box, like picking some classes
    someCourses := courses[1:4] // Get courses from slot 1 to 3 (not including 4)
    fmt.Println("Some courses:", someCourses)
    // This shows: [NodeJS Python Ruby]

    // Loop through courses to see each one, like checking each class
    fmt.Println("Checking each course:")
    for i := 0; i < len(courses); i++ {
        fmt.Println("Slot", i, "has course", courses[i])
    }
    // This shows:
    // Slot 0 has course ReactJS
    // Slot 1 has course NodeJS
    // ...

    // Loop with a magic hand (range) to check courses
    fmt.Println("Using magic hand to check courses:")
    for index, course := range courses {
        fmt.Println("Slot", index, "has course", course)
    }
    // This shows the same courses with their slots

    // Memory tip: Slices share toys underneath
    // Changing someCourses changes courses because they share the same box
    someCourses[0] = "Angular" // Change NodeJS to Angular in the slice
    fmt.Println("Changed some courses:", someCourses)
    fmt.Println("Original courses after change:", courses)
    // This shows Angular in both, like: [ReactJS Angular Python Ruby Course6]

    // Copy courses to a new box to avoid sharing
    coursesCopy := make([]string, len(courses)) // New box same size
    copy(coursesCopy, courses)                 // Copy courses to new box
    fmt.Println("Copied courses:", coursesCopy)
    // This shows: [ReactJS Angular Python Ruby Course6]

    // Change the copy to show it doesn't affect the original
    coursesCopy[0] = "VueJS"
    fmt.Println("Changed copy:", coursesCopy)
    fmt.Println("Original courses unchanged:", courses)
    // This shows copy changed but original stays the same
}