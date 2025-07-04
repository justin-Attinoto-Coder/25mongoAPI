package main

import (
    "fmt"
    "sync"
)

func main() {
    fmt.Println("Race condition - LearnCodeOnline.in")

    wg := &sync.WaitGroup{}
    mut := &sync.Mutex{}
    score := []int{0}

    wg.Add(1)
    go func(wg *sync.WaitGroup, m *sync.Mutex) {
        fmt.Println("One R")
        m.Lock()
        score = append(score, 1)
        m.Unlock()
        wg.Done()
    }(wg, mut)

    wg.Add(1)
    go func(wg *sync.WaitGroup, m *sync.Mutex) {
        fmt.Println("Two R")
        m.Lock()
        score = append(score, 2)
        m.Unlock()
        wg.Done()
    }(wg, mut)

    wg.Add(1)
    go func(wg *sync.WaitGroup, m *sync.Mutex) {
        fmt.Println("Three R")
        m.Lock()
        score = append(score, 3)
        m.Unlock()
        wg.Done()
    }(wg, mut)

    wg.Wait()
    fmt.Println("Final scoreboard:", score)
}

// Five key lessons for a 5-year-old:
// 1. A race condition is when many kids (goroutines) mess up a shared toy box (score) by writing at the same time.
// 2. WaitGroup is like a teacher waiting for all kids to finish their jobs before checking the toy box.
// 3. Mutex is a lock that lets only one kid touch the toy box at a time, keeping it neat.
// 4. Goroutines are like kids working on tasks (adding numbers) all at once to make things faster.
// 5. Using a race detector (go run -race) helps find problems when kids scribble without a lock.