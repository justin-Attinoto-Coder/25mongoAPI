package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    fmt.Println("GoRoutines in golang - LearnCodeOnline.in")

    wg := &sync.WaitGroup{}

    wg.Add(3)
    go func(wg *sync.WaitGroup) {
        for i := 1; i <= 3; i++ {
            fmt.Println("Counting toy", i)
            time.Sleep(time.Millisecond * 500)
        }
        wg.Done()
    }(wg)

    go func(wg *sync.WaitGroup) {
        fmt.Println("Delivering message: Hello, toy shop!")
        fmt.Println("Delivering message: Toys are ready!")
        wg.Done()
    }(wg)

    go func(wg *sync.WaitGroup) {
        fmt.Println("Starting slow task...")
        time.Sleep(time.Second * 2)
        fmt.Println("Slow task done!")
        wg.Done()
    }(wg)

    wg.Wait()
    fmt.Println("All toys counted and tasks done!")
}