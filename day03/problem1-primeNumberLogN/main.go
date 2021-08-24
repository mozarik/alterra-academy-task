package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(PrimeNumber(1000000007))
	fmt.Println(PrimeNumber(1500450271))
	fmt.Println(PrimeNumber(1000000000))
	fmt.Println(PrimeNumber(10000000019))
	fmt.Println(PrimeNumber(10000000033))
}

func PrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	sq_root := int(math.Sqrt(float64(number)))
	for i := 2; i < sq_root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumberNoSquareRoot(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
