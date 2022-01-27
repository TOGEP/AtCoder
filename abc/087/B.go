package main

import "fmt"

func main() {
	var a, b, c, x int
	var cnt = 0

	fmt.Scanf("%d", &a) // 500円
	fmt.Scanf("%d", &b) // 100円
	fmt.Scanf("%d", &c) // 50円
	fmt.Scanf("%d", &x)

	for aCnt := 0; aCnt <= a; aCnt++ {
		for bCnt := 0; bCnt <= b; bCnt++ {
			for cCnt := 0; cCnt <= c; cCnt++ {
				m := 500*aCnt + 100*bCnt + 50*cCnt
				if m == x {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
