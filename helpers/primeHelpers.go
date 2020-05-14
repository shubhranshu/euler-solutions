package helpers

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

//func PrimeProds(primes[] int, num int) []int{
//	prods := make([]int, 0, 100)
//	for idx,prime := range primes{
//
//	}
//}

func checkPrime(primes []int, num int) bool {
	for _, p := range primes {
		if num%p == 0 {
			return false
		}
	}
	return true
}
