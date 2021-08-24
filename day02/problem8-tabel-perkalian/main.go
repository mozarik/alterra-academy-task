package main

import "fmt"

func main() {
	createTable(5)
}

// table perkalian example input 5
// 1	2	3	4	5
// 2	4	6	8	10
// 3	6	9	12	15
// 4	8	12	16	20
// 5	10	15	20	25

func createTable(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}
