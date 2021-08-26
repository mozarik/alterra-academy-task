package main

import (
	"fmt"
	"math"
)

func main() {
	primaSegiEmpat(4, 9, 5)
}

func primaSegiEmpat(height, wide, start int) {
	thePrime := generatePrime(start, height*wide)
	for i := 0; i < height; i++ {
		for j := 0; j < wide; j++ {
			fmt.Print(thePrime[i*wide+j], " ")
		}
		fmt.Println()
	}
}

// Generate prime number from value to n using Sieve of Eratosthenes
func generatePrime(value, n int) (result []int) {

	for i := value; ; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
		if len(result) == n {
			break
		}
	}
	return
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
