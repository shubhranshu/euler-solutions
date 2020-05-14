package main

import (
	"../../helpers"
	"fmt"
	"time"
)

func main() {
	defer helpers.TimeTrack(time.Now(), "Sum 3-5")

	limit := 1000
	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
