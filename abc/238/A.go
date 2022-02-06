package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	// n1 is 2^n
	// n2 is n^2
	n1 := math.Pow(2, float64(n))
	n2 := math.Pow(float64(n), 2)

	if n1 > n2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
