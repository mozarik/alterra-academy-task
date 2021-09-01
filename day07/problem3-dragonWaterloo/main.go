package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	// Search the minimum knight height top chop of the diamer of the dragon head
	var theHeadChopper []int

	// sort both of dragonHead and KnightHeight in ascending order
	sort.Slice(dragonHead, func(i, j int) bool {
		return dragonHead[i] < dragonHead[j]
	})

	sort.Slice(knightHeight, func(i, j int) bool {
		return knightHeight[i] < knightHeight[j]
	})

	// Edge cases
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("Knight Fail")
		return
	}

	// max dragon head
	maxHead := dragonHead[len(dragonHead)-1]
	// min knight height
	maxHeight := knightHeight[len(knightHeight)-1]

	// return knight fail if maxHead > minHeight
	if maxHead > maxHeight {
		fmt.Println("Knight Fail")
		return
	}

	// Check whether the knight can even chop the head of the dragon
	for i := 0; i < len(dragonHead); i++ {
		for j := 0; j < len(knightHeight); j++ {
			if dragonHead[i] <= knightHeight[j] {
				theHeadChopper = append(theHeadChopper, knightHeight[j])
				// Truncate the knightHeight array
				break
			}
		}
	}

	// Sum the head chopper
	fmt.Printf("%v ", theHeadChopper)

	var sum int
	for i := 0; i < len(theHeadChopper); i++ {
		sum += theHeadChopper[i]
	}
	fmt.Println(sum)
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
