/*
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */

package main

import (
	"../../helpers"
	"time"
)

var N = 600851475143

func main() {

	defer helpers.TimeTrack(time.Now(), "Prime factor")

	primes := make([]int, 0, 1000)
	primes = append(primes, 2)
	for i := 2; i < N; i = helpers.NextPrime(primes) {
		for N%i == 0 {
			N = N / i
		}
	}
	println("Answer : ", N)
}
