package main

import "fmt"

func moneyCoins(money int) []int {
	// your code here
	var res []int
	for money != 0 {
		if money >= 10000 {
			res = append(res, 10000)
			money %= 10000
		} else if money >= 5000 {
			res = append(res, 5000)
			money %= 5000
		} else if money >= 2000 {
			res = append(res, 2000)
			money %= 2000
		} else if money >= 1000 {
			res = append(res, 1000)
			money %= 1000
		} else if money >= 500 {
			res = append(res, 500)
			money %= 500
		} else if money >= 200 {
			res = append(res, 200)
			money %= 200
		} else if money >= 100 {
			res = append(res, 100)
			money %= 100
		} else if money >= 50 {
			res = append(res, 50)
			money %= 50
		} else if money >= 20 {
			res = append(res, 20)
			money %= 20
		} else if money >= 10 {
			res = append(res, 10)
			money %= 10
		} else {
			res = append(res, 1)
			money--
		}
	}
	return res
}

func main() {

	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}
