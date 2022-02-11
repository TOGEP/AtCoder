package main

import "fmt"

const N = 3

func sort(num []int) []int {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	return num
}

func main() {
	num := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &num[i])
	}
	num = sort(num)
	fmt.Printf("%d %d %d\n", num[0], num[1], num[2])
}
