package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	t := make([][]int, r+1)
	for i := 0; i < r; i++ {
		t[i] = make([]int, c+1)
		var sum int
		for j := 0; j < c; j++ {
			fmt.Scan(&t[i][j])
			sum += t[i][j]
		}
		t[i][c] = sum
	}

	t[r] = make([]int, c+1)
	for i := 0; i < c+1; i++ {
		var sum int
		for j := 0; j < r; j++ {
			sum += t[j][i]
		}
		t[r][i] = sum
	}

	// print
	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			if j != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", t[i][j])
		}
		fmt.Printf("\n")
	}
	return
}
