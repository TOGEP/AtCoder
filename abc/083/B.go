package main

import "fmt"

func main() {
	var n, a, b int
	var sum = 0

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	num := make([]int, 5)

	for i := 1; i <= n; i++ {
		num[0] = i % 10         // 1の位
		num[1] = i / 10 % 10    // 10の位
		num[2] = i / 100 % 10   // 100の位
		num[3] = i / 1000 % 10  // 1000の位
		num[4] = i / 10000 % 10 // 10000の位
		if num[0]+num[1]+num[2]+num[3]+num[4] >= a && num[0]+num[1]+num[2]+num[3]+num[4] <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}
