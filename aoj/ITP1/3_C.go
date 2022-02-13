package main

import "fmt"

func main() {
	x := make([]int, 3000)
	y := make([]int, 3000)

	for i := 0; ; i++ {
		fmt.Scanf("%d %d", &x[i], &y[i])
		if x[i] == 0 && y[i] == 0 {
			break
		}
		if x[i] > y[i] {
			x[i], y[i] = y[i], x[i]
		}
		fmt.Printf("%d %d\n", x[i], y[i])
	}
}
