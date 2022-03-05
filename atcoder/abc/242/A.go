package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	if x <= a {
		fmt.Printf("%f\n", float64(1.000000000000))
		return
	}
	if x > b {
		fmt.Printf("%f\n", float64(0.000000000000))
		return
	}

	var g float64
	g = float64(c) / (float64(b) - float64(a))
	fmt.Println(g)
	return
}
