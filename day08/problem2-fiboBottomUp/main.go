package main

func main() {

}

// Create a fibonacci function with bottom up approach
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var fibo = make([]int, n+1)
	fibo[0] = 0
	fibo[1] = 1
	for i := 2; i <= n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}
	return fibo[n]
}
