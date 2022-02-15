package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	min := 1000000
	max := -1000000
	sum := 0
	for _, v := range a {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
		sum += v
	}
	fmt.Printf("%d %d %d\n", min, max, sum)
	return
}
