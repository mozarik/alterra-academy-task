package main

import "fmt"

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}

// fast exponentiation function
func pangkat(base, pangkat int) int {
	r := 1
	for pangkat > 0 {
		if pangkat%2 == 1 {
			r *= base
		}
		base *= base
		pangkat /= 2
	}
	return r
}
