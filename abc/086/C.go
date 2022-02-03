package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	t := make([]int, n+1)
	x := make([]int, n+1)
	y := make([]int, n+1)

	// init
	t[0] = 0
	x[0] = 0
	y[0] = 0

	// input
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &t[i+1])
		fmt.Scanf("%d", &x[i+1])
		fmt.Scanf("%d", &y[i+1])
	}

	for i := 0; i < n; i++ {
		//d is distance = |x[i+1] - x[i]| + |y[i+1] - y[i]|
		d := math.Abs(float64(x[i+1]-x[i])) + math.Abs(float64(y[i+1]-y[i]))

		//dt is Δt = t[i+1] - t[i]
		dt := t[i+1] - t[i]

		// MEMO 前提としてdt < t
		if dt < int(d) {
			fmt.Println("No")
			return
		}
		// MEMO その場にとどまってはいけないはdt-dが偶数かどうかで判断できる
		if (dt-int(d))%2 == 1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}
