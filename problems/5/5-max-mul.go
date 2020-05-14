// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"../../helpers"
	"math"
	"time"
)

func main() {
	defer helpers.TimeTrack(time.Now(), "Max prod")
	lim := 20
	primes := make([]int, 0, lim/2)
	primes = append(primes, 2)
	factorMap := make(map[int]int)
	factorMap[2] = 1
	for nextPrime := helpers.NextPrime(primes); nextPrime <= lim; nextPrime = helpers.NextPrime(primes) {
		primes = append(primes, nextPrime)
		factorMap[nextPrime] = 1
	}

	for num := 2; num <= lim; num++ {
		primeLim := int(math.Sqrt(float64(lim)))
		for index := 0; index < len(primes) && primes[index] <= primeLim && num != primes[index]; index++ {
			count := 0
			test := num
			for test%primes[index] == 0 {
				test = test / primes[index]
				count++
			}
			if factorMap[primes[index]] < count {
				factorMap[primes[index]] = count
			}
		}
	}

	prod := 1
	for key, val := range factorMap {
		prod *= int(math.Pow(float64(key), float64(val)))
	}
	println(prod)
}
