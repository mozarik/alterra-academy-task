package main

import "fmt"

func main() {
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Printf("a : %d, b: %d", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
