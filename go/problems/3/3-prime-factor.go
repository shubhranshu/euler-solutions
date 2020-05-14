/*
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */

package main

var N = 600851475143
var primes = make([]int, 0, 1000)

func main() {

	primes = append(primes, 2)
	for i := 2; i < N; i = nextPrime() {
		for N%i == 0 {
			N = N / i
		}
	}
	println("Answer : ", N)
}

func nextPrime() int {
	lastPrime := primes[len(primes)-1]
	candidate := lastPrime + 1
	var newPrime int
	for newPrime == 0 {
		if checkPrime(candidate) {
			primes = append(primes, candidate)
			newPrime = candidate
		}
		candidate = candidate + 1
	}
	return newPrime
}

func checkPrime(num int) bool {
	for _, p := range primes {
		if num%p == 0 {
			return false
		}
	}
	return true
}
