package main

import "fmt"

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}

// This function will take a slice of 2d arrays and return true of the we have same card on deck
func playingDomino(cards [][]int, deck []int) interface{} {
	for i := 0; i < len(cards); i++ {
		if cards[i][0] == deck[0] || cards[i][1] == deck[0] {
			return true
		}
	}
	return false
}
