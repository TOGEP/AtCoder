package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// 長方形の面積S=a*b
	s := a * b

	// 長方形の周の長さ L=(a+b)*2
	l := (a + b) * 2

	fmt.Printf("%d %d\n", s, l)
}
