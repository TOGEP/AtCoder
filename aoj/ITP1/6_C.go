package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var s [4][3][10]int
	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scanf("%d %d %d %d", &b, &f, &r, &v)
		s[b-1][f-1][r-1] += v
	}

	for i := range []int{0, 1, 2, 3} {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", s[i][j][k])
			}
			fmt.Printf("\n")
		}
		if i != 3 {
			fmt.Println("####################")
		}
	}
	return
}
