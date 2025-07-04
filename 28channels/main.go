package main

import (
    "fmt"
    "sync"
)

func main() {
    fmt.Println("Channels in golang - LearnCodeOnline.in")

    myCh := make(chan int)
    wg := &sync.WaitGroup{}

    wg.Add(3)
    go func(ch chan int, wg *sync.WaitGroup) {
        ch <- 5
        ch <- 10
        wg.Done()
    }(myCh, wg)

    go func(ch chan int, wg *sync.WaitGroup) {
        close(ch)
        wg.Done()
    }(myCh, wg)

    go func(ch chan int, wg *sync.WaitGroup) {
        val, isChannelOpen := <-myCh
        fmt.Println("Value:", val, "Is channel open?", isChannelOpen)
        val, isChannelOpen = <-myCh
        fmt.Println("Value:", val, "Is channel open?", isChannelOpen)
        wg.Done()
    }(myCh, wg)

    wg.Wait()
}