package main

import (
    "fmt"
    "os"
    "runtime"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our building adventure
    fmt.Println("Welcome to building Go programs for different computers!")

    // Let's find out what kind of computer we're using right now
    // GOHOSTOS tells us the operating system of our own computer
    hostOS := runtime.GOOS // This is like asking, "What kind of toy box are we in?"
    fmt.Println("Our computer's operating system (GOHOSTOS):", hostOS)
    // This might print: windows, linux, or darwin (Mac)

    // We can also check GOOS, which is what we want to build for
    // GOOS is set when we build, but here we just show the default
    fmt.Println("Default target operating system (GOOS):", runtime.GOOS)
    // This usually matches GOHOSTOS unless we change it

    // Let's make a simple program that says who we are
    user := "Justin"
    fmt.Println("Hello from", user, "on", hostOS, "!")

    // Check if we're running on the right system
    if hostOS == "windows" {
        fmt.Println("Yay, we're on Windows! This program loves .exe files!")
    } else if hostOS == "linux" {
        fmt.Println("Cool, we're on Linux! This program runs without extensions!")
    } else if hostOS == "darwin" {
        fmt.Println("Awesome, we're on a Mac! This program loves macOS!")
    }

    // Print some Go environment info to understand our setup
    // GOARCH is the type of computer brain (like 64-bit or 32-bit)
    goArch := os.Getenv("GOARCH")
    fmt.Println("Our computer's brain type (GOARCH):", goArch)
    // This might print: amd64 (most computers) or arm64 (some Macs)

    // GOPATH is where Go keeps its toys (code and tools)
    goPath := os.Getenv("GOPATH")
    fmt.Println("Where Go keeps its toys (GOPATH):", goPath)
    // This might print: /home/justin/go or C:\Users\Justin\go
}