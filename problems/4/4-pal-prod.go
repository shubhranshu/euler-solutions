package main

import (
	"../../helpers"
	"math"
	"strconv"
	"time"
)

const mulDigits = 3

var testCount = 0

func main() {
	defer helpers.TimeTrack(time.Now(), "Pal Prod")
	max := int(math.Pow10(mulDigits) - 1)
	min := int(math.Pow10(mulDigits - 1))
	println("min max limits ", min, max)
	println("MaxProd: ", maxProd(min, max))
	println("Test count : ", testCount)
}

func maxProd(min int, max int) int {
	candidateProd := 0
	di := 1
	dj := 1
	j := max
	for i := max; i >= min; i -= di {
		if i%11 == 0 {
			j = max
			dj = 1
		} else {
			j = 990
			dj = 11
		}
		for ; j >= i; j -= dj {
			prod := i * j
			if prod <= candidateProd {
				break
			}
			if isPalindrome(prod) {
				candidateProd = prod
			}
		}
	}
	return candidateProd
}

func isPalindrome(prod int) bool {
	testCount++
	strNum := strconv.Itoa(prod)
	for i := 0; i < len(strNum)/2; i++ {
		if strNum[i] != strNum[len(strNum)-i-1] {
			return false
		}
	}
	return true
}
