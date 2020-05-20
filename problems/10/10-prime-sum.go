/*

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/

package main

import (
	"../../helpers"
	"time"
)

func main() {
	defer helpers.TimeTrack(time.Now(), "Max prod")

	primes := helpers.BasicPrimes(2000000)

	sum := 0
	for _, prime := range primes {
		sum += prime
	}

	println(sum)

}
