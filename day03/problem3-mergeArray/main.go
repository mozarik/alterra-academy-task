package main

import "fmt"

// merge 2 arrays without duplicate value
func ArrayMerge(a, b []string) []string {
	var result []string

	// if a and b is empty, return a
	if len(a) == 0 && len(b) == 0 {
		return a
	}

	// if a is empty, return b
	if len(a) == 0 {
		return b
	}

	// if b is empty, return a
	if len(b) == 0 {
		return a
	}

	// if a and b has same value, return a
	// array A always in the beginning of the output so we append a to result directly
	result = append(result, a...)

	// In here we check if b is in result, if not, append it to result
	for _, v := range b {
		if !contains(result, v) {
			result = append(result, v)
		}
	}

	return result
}

// Check if array contains value
func contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
