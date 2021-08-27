package main

import "fmt"

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// Create a MaxSequence function that that return the sum of the maximum sequence of consecutive integers in an array.
// The maximum sequence is defined as the longest sequence of integers in the array that have a sum equal to the maximum sum in the array.
// For example, the maximum sequence of consecutive integers in the array {1, 2, -3, 4, -1, 2, 1, -5, 4} is {4, -1, 2, 1} with a sum of 6.
// The array {1, 2, -3, 4, -1, 2, 1, -5, 4} has a maximum sum of 6 and a maximum sequence of {4, -1, 2, 1}.
// The array {-2, 1, -3, 4, -1, 2, 1, -5, 4} has a maximum sum of 7 and a maximum sequence of {1, -5, 4}.
// The array {-2, 1, -3, 4, -1, 2, 1, -5, 4, -1} has a maximum sum of 8 and a maximum sequence of {4, -1, 2, 1}.
// The array {-2, 1, -3, 4, -1, 2, 1, -5, 4, -1, 6} has a maximum sum of 13 and a maximum sequence of {4, -1, 2, 1, 6}.
// The array {-2, 1, -3, 4, -1, 2, 1, -5, 4, -1, 6, -7, -4, -8, -1, -2, -9} has a maximum sum of 18 and a maximum sequence of {4, -1, 2, 1, 6, -7, -4, -8, -1, -2, -9}.

func MaxSequence(arr []int) int {
	result := MaxSequenceValue(MaxSequenceArray(arr))
	return result
}

func MaxSequenceArray(arr []int) []int {
	var maxSum int
	var maxSequence []int
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
				maxSequence = arr[i : j+1]
			}
		}
	}
	return maxSequence
}

func MaxSequenceValue(arr []int) int {
	resultArray := MaxSequenceArray(arr)
	var sum int
	for i := 0; i < len(resultArray); i++ {
		sum += resultArray[i]
	}
	return sum
}
