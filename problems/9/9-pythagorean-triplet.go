/*

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

166167000
Prod :  31875000
Max prod : 162.5655ms



*/

package main

import (
	"../../helpers"
	"math"
	"time"
)

func main() {
	defer helpers.TimeTrack(time.Now(), "Max prod")

	prod := 0
	s:=1000
	s2 := s/2
	limit := int(math.Ceil(math.Sqrt(float64(s2)))) -1

	for m := 2; m < limit; m++ {
		if s2 %m == 0 {
			sm := s2/m
			for sm %2 == 0{
				sm = sm/2
			}
			var k int
			if m%2 == 1 {
				k = m+2
			}else{
				k = m+1
			}
			for k< 2*m && k<= sm{
				if sm % k == 0 && helpers.GCD(k,m) ==1 {
					d := s2/(k*m)
					n := k-m
					a := d*(m*m-n*n)
					b := 2*d*m*n
					c := d*(m*m+n*n)
					prod = a*b*c
					goto done
				}else{
					k += 2
				}
			}
		}
	}

	done:
	println("Prod : ", prod)

}
