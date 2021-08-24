package main

import "fmt"

func main() {
	var bilagan int
	fmt.Scanf("%d", &bilagan)

	for _, v := range factor(bilagan) {
		fmt.Println(v)
	}
}

// Function to find factor of a number
func factor(num int) []int {
	var factors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
