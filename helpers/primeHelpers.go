package helpers

import "math"

func NextPrime(primes []int) int {
	lastPrime := primes[len(primes)-1]
	candidate := lastPrime + 1
	var newPrime int
	for newPrime == 0 {
		if checkPrime(primes, candidate) {
			primes = append(primes, candidate)
			newPrime = candidate
		}
		candidate = candidate + 1
	}
	return newPrime
}

func checkPrime(primes []int, num int) bool {
	for _, p := range primes {
		if num%p == 0 {
			return false
		}
	}
	return true
}

func PrimeAtkins(primeLimit int) []int {

	sqLimit := int(math.Sqrt(float64(primeLimit)))
	primeSieve := make([]bool, primeLimit, primeLimit)

	for x := 1; x <= sqLimit/2; x++ {
		for y := 1; y < sqLimit; y += 2 {
			n := 4*x*x + y*y
			if n >= primeLimit {
				continue
			}
			mod := n % 60
			if mod == 1 || mod == 13 || mod == 17 || mod == 29 || mod == 37 || mod == 41 || mod == 49 || mod == 53 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for x := 1; x < int(float64(sqLimit)/1.6); x += 2 {
		for y := 2; y < sqLimit; y += 2 {
			n := 3*x*x + y*y
			if n >= primeLimit {
				continue
			}
			mod := n % 60
			if mod == 7 || mod == 19 || mod == 31 || mod == 43 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for x := 2; x < sqLimit; x += 1 {
		for y := 1; y < sqLimit && y <= x; y += 2 {
			n := 3*x*x - (x-y)*(x-y)
			if n >= primeLimit {
				continue
			}

			mod := n % 60
			if mod == 11 || mod == 23 || mod == 47 || mod == 59 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for n := 0; n < sqLimit; n++ {
		if primeSieve[n] {
			for y := n * n; y < primeLimit; y += n * n {
				primeSieve[y] = false
			}
		}
	}

	// optimistically allocate memory
	primes := make([]int, 0, primeLimit/20)
	primes = append(primes, 2)
	primes = append(primes, 3)
	primes = append(primes, 5)
	for index, prime := range primeSieve {
		if prime {
			primes = append(primes, index)
		}
	}

	return primes
}

func BasicPrimes(primeLimit int) []int {

	primes := make([]int, 0, primeLimit/10)

	primes = append(primes, 2)

	for i := 3; i < primeLimit; i = i + 2 {
		sqPrime := int(math.Sqrt(float64(i)))
		for _, prime := range primes {
			if prime > sqPrime {
				primes = append(primes, i)
				break
			}
			if i%prime == 0 { //not prime
				break
			}
		}
	}

	return primes
}
