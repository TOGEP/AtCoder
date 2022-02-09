package main

import "fmt"

func main() {
	var h, w int

	fmt.Scanf("%d %d", &h, &w)
	matrix := make([][]int, h)
	for row := range matrix {
		matrix[row] = make([]int, w)
		for column := range matrix[row] {
			fmt.Scanf("%d", &matrix[row][column])
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf("%d ", matrix[j][i])
		}
		fmt.Printf("\n")
	}
	return
}
