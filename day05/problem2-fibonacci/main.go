package main

import "fmt"

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(12)) // 144
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
