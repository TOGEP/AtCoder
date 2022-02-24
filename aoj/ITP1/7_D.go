package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
		for j := 0; j < l; j++ {
			var sum int
			for k := 0; k < m; k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}

	// print
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			if j != 0 {
				fmt.Printf(" ")
			}
			fmt.Print(c[i][j])
		}
		fmt.Printf("\n")
	}
	return
}
