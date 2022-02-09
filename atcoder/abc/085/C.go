package main

import (
	"fmt"
	"math"
)

// N次連立方程式
const N = 2

// プログラムが簡単，制度の高さからガウス・ジョルダン法を採用
func GaussJordan(x [N][N + 1]int) [N]int {
	var ret [N]int
	var pivot int
	var eps = 0.0001 // 微小量の設定

	for i := 0; i < N; i++ {
		pivot = x[i][i]
		if math.Abs(float64(pivot)) < eps {
			for j := range ret {
				ret[j] = -1
			}
			return ret
		}
		for j := 0; j < N+1; j++ {
			x[i][j] = x[i][j] / pivot
		}

		for k := 0; k < N; k++ {
			del := x[k][i]
			for j := i; j < N+1; j++ {
				if k != i {
					x[k][j] = x[k][j] - del*x[i][j]
				}
			}
		}
	}

	// DEBUG print
	/*
		for i := 0; i < N; i++ {
			for j := 0; j < N+1; j++ {
				fmt.Printf("%d ", x[i][j])
			}
			fmt.Printf("\n")
		}
	*/

	for i := 0; i < N; i++ {
		ret[i] = x[i][N]
	}
	return ret
}

func main() {
	var n, y int
	// ret[0], ret[1], ret[2] is number of 10000, 5000, 1000 yen
	var ret [N]int

	fmt.Scanf("%d %d", &n, &y)

	// 以下の連立方程式を考える
	// a + b = n
	// 10000a + 5000b = y
	x := [N][N + 1]int{
		{1, 1, n},
		{10000, 5000, y},
	}

	ret = GaussJordan(x)
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", ret[i])
	}
	// c = N - (a+b)
	fmt.Printf("%d\n", n-(ret[0]+ret[1]))

	return
}
