package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var d int = a / b
	var r int = a % b
	var f float64 = float64(a) / float64(b)

	fmt.Printf("%d %d %f", d, r, f)
}
