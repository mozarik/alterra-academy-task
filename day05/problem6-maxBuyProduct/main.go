package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(MaxBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
}

// MaxBuyProduct based on the given input, find the maximum number of products that can be bought
// Input is the money we have and list of product price

func MaxBuyProduct(money int, price []int) int {
	var count int
	var sum int

	sort.IntSlice.Sort(price)
	for i := 0; i < len(price); i++ {
		if sum+price[i] <= money {
			sum += price[i]
			count++
		}
	}

	return count
}
