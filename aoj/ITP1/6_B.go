package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ss := make([][]int, 4)
	for i := 0; i < 4; i++ {
		ss[i] = make([]int, 13)
	}
	for i := 0; i < n; i++ {
		var k string
		var v int
		fmt.Scanf("%s %d", &k, &v)
		v = v - 1
		switch k {
		// MEMO intの初期値は0なので，ss[key][v] = vにすると1が絶対に見つかる判定となる
		case "S":
			ss[0][v] = 1
		case "H":
			ss[1][v] = 1
		case "C":
			ss[2][v] = 1
		case "D":
			ss[3][v] = 1
		}
	}

	for i, v := range []string{"S", "H", "C", "D"} {
		for j := 0; j < 13; j++ {
			if ss[i][j] != 1 {
				fmt.Printf("%s %d\n", v, j+1)
			}
		}
	}
	return
}
