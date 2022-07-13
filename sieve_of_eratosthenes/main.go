package main

import "fmt"

func main() {
	fmt.Println(findAllPrimes(50)) // [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
}

//findAllPrimes finds all the prime numbers less than or equal to a given integer n by the Eratosthenes’s method
func findAllPrimes(n int) []int {
	allPrimes := make([]bool, n+1)
	var result []int

	if n < 2 {
		return result
	}

	// initialize all entries it as true. A value in prime[i] will finally be false if i is Not a prime, else true.
	for i := range allPrimes {
		allPrimes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If prime[p] is not changed(true), then it is a prime
		if allPrimes[p] {
			// Update all multiples of p greater than or equal to the square of it numbers which are multiple of p
			//and are less than p^2 are already been marked.
			for j := p * p; j <= n; j += p {
				allPrimes[j] = false
			}
		}
	}

	// print all primes
	for i, v := range allPrimes {
		if i < 2 {
			continue
		}
		if v {
			result = append(result, i)
		}
	}
	return result

}
