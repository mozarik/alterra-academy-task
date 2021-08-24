package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}

// Pairsum returns a slice of the indices of the two numbers in the
// given array that sum to the target.
// return the index of the pair of numbers that sum to the target
// or return an empty slice if no pair exists
func PairSum(a []int, target int) []int {
	// use a map to store index of each number

	indexValue := map[int]int{}
	for i, v := range a {
		indexValue[v] = i
	}

	set := mapset.NewSet()
	for i := 0; i < len(a); i++ {
		temp := target - a[i]
		if set.Contains(temp) {
			indexOfTemp := indexValue[temp]
			return []int{indexOfTemp, i}
		}
		set.Add(a[i])
	}
	return []int{}
}
