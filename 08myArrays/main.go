package main

import "fmt"

// The main function is like the start button for our program!
func main() {
    // Say hello to our array adventure
    fmt.Println("Welcome to our fruit basket adventure with arrays!")

    // An array is like a toy box with a fixed number of slots for toys
    // Let's make a box with 4 slots for fruit names
    var fruitBasket [4]string
    // This makes a box with 4 empty slots: ["", "", "", ""]

    // Put fruits in the slots, like putting toys in a box
    fruitBasket[0] = "Apple"   // Slot 1 gets an Apple
    fruitBasket[1] = "Banana"  // Slot 2 gets a Banana
    fruitBasket[2] = "Orange"  // Slot 3 gets an Orange
    fruitBasket[3] = "Mango"   // Slot 4 gets a Mango

    // Show the whole fruit basket
    fmt.Println("My fruit basket:", fruitBasket)
    // This shows: [Apple Banana Orange Mango]

    // Show just one fruit by looking in a slot
    fmt.Println("First fruit in slot 0:", fruitBasket[0])
    // This shows: Apple

    // Change a fruit in a slot, like swapping a toy
    fruitBasket[2] = "Grape" // Replace Orange with Grape
    fmt.Println("New fruit basket:", fruitBasket)
    // This shows: [Apple Banana Grape Mango]

    // Find out how many slots the basket has
    basketSize := len(fruitBasket)
    fmt.Println("My basket has", basketSize, "slots")
    // This shows: My basket has 4 slots

    // Make another fruit basket with fruits already in it
    veggieBasket := [3]string{"Carrot", "Potato", "Cucumber"}
    fmt.Println("My veggie basket:", veggieBasket)
    // This shows: [Carrot Potato Cucumber]

    // Loop through the basket to see each fruit, like checking each slot
    fmt.Println("Checking each fruit in the fruit basket:")
    for i := 0; i < len(fruitBasket); i++ {
        fmt.Println("Slot", i, "has", fruitBasket[i])
    }
    // This shows:
    // Slot 0 has Apple
    // Slot 1 has Banana
    // Slot 2 has Grape
    // Slot 3 has Mango

    // Another way to loop using range, like a magic hand picking each fruit
    fmt.Println("Using magic hand to check veggies:")
    for index, veggie := range veggieBasket {
        fmt.Println("Slot", index, "has", veggie)
    }
    // This shows:
    // Slot 0 has Carrot
    // Slot 1 has Potato
    // Slot 2 has Cucumber

    // Make a big basket with numbers, like counting fruits
    fruitCount := [5]int{10, 20, 15, 30, 25} // Number of each fruit
    fmt.Println("Fruit counts:", fruitCount)
    // This shows: [10 20 15 30 25]

    // Add up all the fruits, like counting all toys
    totalFruits := 0
    for _, count := range fruitCount {
        totalFruits = totalFruits + count
    }
    fmt.Println("Total fruits in the big basket:", totalFruits)
    // This shows: 100 (10 + 20 + 15 + 30 + 25)

    // Make a double basket, like a box with rows and columns for fruits
    // This is a 2x3 box (2 rows, 3 columns)
    doubleBasket := [2][3]string{
        {"Apple", "Banana", "Grape"},    // Row 1
        {"Mango", "Orange", "Pineapple"}, // Row 2
    }
    fmt.Println("My double basket:", doubleBasket)
    // This shows: [[Apple Banana Grape] [Mango Orange Pineapple]]

    // Look at one fruit in the double basket
    fmt.Println("Fruit in row 1, column 2:", doubleBasket[0][1])
    // This shows: Banana (row 1 starts at 0, column 2 starts at 1)

    // Loop through the double basket to see all fruits
    fmt.Println("Checking all fruits in the double basket:")
    for row := 0; row < len(doubleBasket); row++ {
        for col := 0; col < len(doubleBasket[row]); col++ {
            fmt.Println("Row", row, "Column", col, "has", doubleBasket[row][col])
        }
    }
    // This shows:
    // Row 0 Column 0 has Apple
    // Row 0 Column 1 has Banana
    // Row 0 Column 2 has Grape
    // Row 1 Column 0 has Mango
    // Row 1 Column 1 has Orange
    // Row 1 Column 2 has Pineapple
}