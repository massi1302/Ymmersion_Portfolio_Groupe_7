package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    
    targetNumber := rand.Intn(99) + 1
    
    var guess int
    guessCount := 0
    
    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I've chosen a number between 1 and 99. Try to guess it!")
    
    for {
        fmt.Print("Enter your guess: ")
        fmt.Scan(&guess)
        guessCount++
        
        if guess < targetNumber {
            fmt.Println("Too low! The number is higher.")
        } else if guess > targetNumber {
            fmt.Println("Too high! The number is lower.")
        } else {
            fmt.Printf("Congratulations! You guessed the number %d in %d guesses!\n", targetNumber, guessCount)
            break
        }
    }
}