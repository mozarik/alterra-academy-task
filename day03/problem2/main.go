package main

import "fmt"

func main() {
	fmt.Print(pow(2, 50))
	fmt.Print(pow(7, 2))
	fmt.Print(pow(10, 5))
	fmt.Print(pow(17, 6))
	fmt.Print(pow(5, 3))
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
