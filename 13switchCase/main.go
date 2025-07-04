package main

import (
    "fmt"
    "math/rand"
    "time"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our switch case adventure
    fmt.Println("Welcome to our fruit board dice game adventure!")

    // Set up the dice roller to pick different numbers each time
    rand.Seed(time.Now().UnixNano()) // Like mixing up a bag of dice

    // Our game board has fruit spots, like a path to walk
    board := []string{"Start", "Apple", "Banana", "Orange", "Mango", "Peach", "Grape", "Finish"}
    playerPosition := 0 // Player starts at the "Start" spot
    maxPosition := len(board) - 1 // The "Finish" spot

    // Keep playing until the player reaches or passes the Finish
    fmt.Println("Let's roll the dice and move on the fruit board!")
    for playerPosition < maxPosition {
        // Roll the dice (1 to 6)
        diceNumber := rand.Intn(6) + 1 // Pick a random number from 1 to 6
        fmt.Println("Value of dice is", diceNumber)

        // Use a switch to decide what happens based on the dice
        switch diceNumber {
        case 1:
            // Move 1 spot, like a small step
            playerPosition += 1
            fmt.Println("You move 1 spot!")
        case 2:
            // Move 2 spots, like a bigger step
            playerPosition += 2
            fmt.Println("You move 2 spots!")
        case 3:
            // Move 3 spots
            playerPosition += 3
            fmt.Println("You move 3 spots!")
        case 4:
            // Move 4 spots
            playerPosition += 4
            fmt.Println("You move 4 spots!")
        case 5:
            // Move 5 spots
            playerPosition += 5
            fmt.Println("You move 5 spots!")
        case 6:
            // Move 6 spots and roll again!
            playerPosition += 6
            fmt.Println("Wow, a 6! You move 6 spots and roll again!")
            continue // Skip the rest and roll a new dice
        }

        // Make sure the player doesn't go past the Finish
        if playerPosition > maxPosition {
            playerPosition = maxPosition
        }

        // Show where the player is on the fruit board
        fmt.Println("You're now at", board[playerPosition])
        fmt.Println("---")
    }

    // Check if the player won by reaching the Finish
    if playerPosition == maxPosition {
        fmt.Println("Yay! You reached the Finish! You win the fruit game!")
    }

    // Try another roll to show a different switch style
    fmt.Println("\nLet's roll one more dice for fun!")
    anotherDice := rand.Intn(6) + 1
    fmt.Println("Dice value:", anotherDice)
    switch {
    case anotherDice <= 3:
        fmt.Println("Low roll! You get a small fruit prize!")
    case anotherDice >= 4:
        fmt.Println("High roll! You get a big fruit basket!")
    }
}

// Explanation comments for a 5-year-old:
// - What's a switch statement? It's like a toy sorter that looks at a toy (dice number) and picks a path to follow, like "If it's a 6, do this!"
// - What's 'switch diceNumber'? It's like asking, "What number is on the dice?" and checking each possible number (1, 2, 3, 4, 5, 6).
// - What's 'case 1'? It's like saying, "If the dice is 1, move 1 spot on the board!"
// - What's 'case 6: ... continue'? It's like saying, "If the dice is 6, move 6 spots and roll again!" The 'continue' makes us skip to a new roll.
// - What's 'rand.Seed(time.Now().UnixNano())'? It's like mixing up a bag of dice so we get different numbers each time we play.
// - What's 'diceNumber := rand.Intn(6) + 1'? It's like picking a random dice from 1 to 6, like rolling a real dice.
// - What's 'board := []string{"Start", "Apple", ...}'? It's like a path with fruit names as spots, like stepping stones in a game.
// - What's 'playerPosition += 1'? It's like moving the player 1 step forward on the fruit board.
// - What's 'if playerPosition > maxPosition'? It's like making sure the player stops at the "Finish" spot and doesn't go too far.
// - What's 'switch { case anotherDice <= 3: ... }'? It's another way to sort the dice, like saying, "If the number is 3 or less, get a small prize!"
// - Why a dice game? It's fun and shows how switch picks different actions for each dice number, like moving or rolling again.
// - Why fruit board? It keeps the fruit theme (like Apple, Mango) from earlier chapters to make the game feel familiar.
// - How does this help beginners? It shows how switch makes decisions easier than lots of if-else, with a fun game where you move on a fruit board and roll again on 6.