package main

import "fmt"

func main() {
	playWithAsterik(5)
}

// Function to create string of playWithAsterik pyramide
func playWithAsterik(n int) {
	spaces := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		spaces = spaces - 1

		for k := 0; k < i+1; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
