package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generates random number from 1 to 25
	target := rand.Intn(25) + 1

	fmt.Println("Welcome to the game 'Guess the number'!")
	fmt.Println("I thought in number from 1 to 25. Try to guess!")

	reader := bufio.NewReader(os.Stdin)

	attempts := 0
	for {
		attempts++
		fmt.Print("Type your number: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}

		// Удаляем все пробельные символы, включая '\r' и '\n'
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Try again")
			continue
		}

		if guess < target {
			fmt.Println("The number is bigger")
		} else if guess > target {
			fmt.Println("The number is smaller")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts!\n", attempts)
			break
		}
	}
}
