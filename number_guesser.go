package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var solution int = rand.Intn(101)
	fmt.Println("Guess a number between 0 and 100")

	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		_ = err
		var output string = guess_checker(guess, solution)
		fmt.Println(output)
		if output == "You win!" {
			break
		}
	}
}

func guess_checker(guess int, solution int) string {
	if guess == solution {
		return "You win!"
	} else if guess < solution {
		return "Too low"
	} else {
		return "Too high"
	}
}
