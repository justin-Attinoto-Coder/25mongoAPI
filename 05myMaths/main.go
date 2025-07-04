package main

import (
    cryptoRand "crypto/rand"
    "fmt"
    "math"
    "math/big"
    mathRand "math/rand"
    "time"
)

func main() {
    fmt.Println("Welcome to our math playground!")

    number1 := 10.5
    number2 := 3

    sum := number1 + float64(number2)
    fmt.Printf("Adding %v and %v gives us %v!\n", number1, number2, sum)

    difference := number1 - float64(number2)
    fmt.Printf("Taking %v away from %v leaves %v!\n", number2, number1, difference)

    product := number1 * float64(number2)
    fmt.Printf("%v groups of %v makes %v!\n", number2, number1, product)

    quotient := number1 / float64(number2)
    fmt.Printf("Sharing %v among %v gives %v each!\n", number1, number2, quotient)

    rounded := math.Round(number1)
    fmt.Printf("Rounding %v gives us %v!\n", number1, rounded)

    ceiling := math.Ceil(number1)
    fmt.Printf("Ceiling of %v is %v!\n", number1, ceiling)

    floor := math.Floor(number1)
    fmt.Printf("Floor of %v is %v!\n", number1, floor)

    maxNumber := math.Max(number1, float64(number2))
    fmt.Printf("The bigger number between %v and %v is %v!\n", number1, number2, maxNumber)

    power := math.Pow(2, 3)
    fmt.Printf("2 grown 3 times is %v!\n", power)

    sqrt := math.Sqrt(16)
    fmt.Printf("The square root of 16 is %v!\n", sqrt)

    mathRand.Seed(time.Now().UnixNano())
    fmt.Println("Random number:", mathRand.Intn(5))

    myRandomNum, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(5))
    fmt.Println("Crypto random number:", myRandomNum)
}

// Five key lessons for a 5-year-old:
// 1. Math operations like adding and sharing toys help us play with numbers.
// 2. The math package is a toy box with tools like rounding and square roots.
// 3. Random numbers are like picking a surprise toy from a box.
// 4. Crypto random numbers are super safe surprises for important games.
// 5. Import aliases keep different toy boxes from getting mixed up.