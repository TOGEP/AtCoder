package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	var cnt int
	for cnt = 0; ; cnt++ {
		for i := 0; i < N; i++ {
			if a[i]%2 == 1 || a[i] == 0 {
				fmt.Printf("%d\n", cnt)
				return
			}
			a[i] = a[i] / 2
		}
	}
}
