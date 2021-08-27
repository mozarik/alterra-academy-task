package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindMinAndMAx([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(FindMinAndMAx([]int{-2, 5, -7, 4, 7, -20}))
}

// Return Min Value and the index and Max Value and the Index
func FindMinAndMAx(arr []int) string {
	min := arr[0]
	max := arr[0]
	minIndex := 0
	maxIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}
	return "Min Value: " + strconv.Itoa(min) + " at Index: " + strconv.Itoa(minIndex) + " Max Value: " + strconv.Itoa(max) + " at Index: " + strconv.Itoa(maxIndex)
}
