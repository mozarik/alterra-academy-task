package main

// Create sieve of Eratosthenes

func main() {
	var n int
	var primes []int

	// Get the number
	n = 1000000

	// Get the primes
	primes = sieveOfEratosthenes(n)

	// Print the primes
	for i := 0; i < len(primes); i++ {
		println(primes[i])
	}
}

// Sieve of Eratosthenes function
func sieveOfEratosthenes(n int) []int {
	var primes []int
	var isPrime []bool
	var i int
	var j int

	isPrime = make([]bool, n+1)

	// There is no magic in Go. so we have to do it manually, lol no memset
	for i = 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Cross out non-primes multiples
	for i = 2; i*i <= n; i++ {
		if isPrime[i] == true {
			for j = i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Add primes to the list
	for i = 2; i <= n; i++ {
		if isPrime[i] == true {
			primes = append(primes, i)
		}
	}

	return primes
}
