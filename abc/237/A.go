package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	fmt.Scanf("%d", &n)

	min := math.Pow(-2, 31)
	max := math.Pow(2, 32)
	if min <= float64(n) && max > float64(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
