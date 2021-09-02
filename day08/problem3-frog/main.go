package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
}

// Find rthe minimum possible total cost incurred before the frog reacher stone N
func Frog(jumps []int) int {
	var dp []int
	n := len(jumps)
	var h []int

	for i := 0; i < n; i++ {
		h = append(h, jumps[i])
	}

	dp = append(dp, 0)

	for i := 1; i < n; i++ {
		if i >= 2 {
			var tmp1, tmp2 int
			tmp1 = dp[i-1] + int(math.Abs(float64(h[i]-h[i-1])))
			tmp2 = dp[i-2] + int(math.Abs(float64(h[i]-h[i-2])))

			if tmp1 < tmp2 {
				dp = append(dp, tmp1)
			} else {
				dp = append(dp, tmp2)
			}
		} else {
			dp = append(dp, dp[i-1]+int(math.Abs(float64(h[i]-h[i-1]))))
		}
	}
	return dp[len(dp)-1]
}
