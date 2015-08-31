package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {

	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	// while
	sum2 := 1
	for sum2 < 1000 {
		sum += sum2
	}
	fmt.Println(sum2)

	// infinity loop
	//	for {
	//
	//	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}
