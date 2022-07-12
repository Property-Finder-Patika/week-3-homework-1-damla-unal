package main

import "fmt"

func main() {

	fmt.Println("Prime factors for 12: ", getPrimeFactors(12))
}

func getPrimeFactors(n int) []int {
	var primeFactors []int
	//find all prime factor 2 so repeatedly divide by 2
	for n%2 == 0 {
		primeFactors = append(primeFactors, 2)
		n /= 2
	}

	//start with prime number 3 and increment each by two to skip even number
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
	}

	//check n is a prime number
	if n > 2 {
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}
