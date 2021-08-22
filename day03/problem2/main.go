package main

import "fmt"

func main() {
	fmt.Print(pow(2, 50))
}

// Create fast exponentiation function
func pow(x, n int) int {
	r := 1
	for n > 0 {
		if n%2 == 1 {
			r *= x
		}
		x *= x
		n /= 2
	}
	return r
}
