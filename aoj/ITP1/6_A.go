package main

import "fmt"

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	reverse(a)

	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
	return
}
