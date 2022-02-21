package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	dp := make([][]int, 101)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 10001)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= x; j++ {
			if dp[i][j] == 1 {
				if j+a[i] <= x {
					dp[i+1][j+a[i]] = 1
				}
				if j+b[i] <= x {
					dp[i+1][j+b[i]] = 1
				}
			}
		}
	}
	if dp[n][x] == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
