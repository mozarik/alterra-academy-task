package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// your code here
	for i := 0; i <= 600; i++ {
		for j := 0; j <= 600; j++ {
			for k := 0; k <= 600; k++ {
				z := int(math.Pow(float64(i), 2)) + int(math.Pow(float64(j), 2)) + int(math.Pow(float64(k), 2))
				if i+j+k == a && i*j*k == b && z == c {
					fmt.Println(i, j, k)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {

	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3

}
