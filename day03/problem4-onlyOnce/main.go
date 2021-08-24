package main

import (
	"fmt"
	"strconv"
)

// function to list number that only appear once in a string
func munculSekali(numbers string) []int {
	var seen = make(map[string]string)
	var stringNumValue string

	// Input first number to map
	stringNumValue = string(numbers[0])
	seen[stringNumValue] = stringNumValue

	// Loop through each number in the string
	for i := 1; i < len(numbers); i++ {
		stringNumValue = string(numbers[i])
		if _, ok := seen[stringNumValue]; ok {
			delete(seen, stringNumValue)
		} else {
			seen[stringNumValue] = stringNumValue
		}
	}

	// Convert map to slice
	var result []int
	for key, _ := range seen {
		keyInt, _ := strconv.Atoi(key)
		result = append(result, keyInt)
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
