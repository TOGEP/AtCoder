package main

import (
	"fmt"
)

func remove(ints []int, search int) []int {
	ret := []int{}
	for _, v := range ints {
		if v != search {
			ret = append(ret, v)
		}
	}
	return ret
}

func generateX(x []int, num int) []int {
	x = remove(x, num)
	x1 := num / 2
	x2 := num - x1
	x = append(x, x1)
	x = append(x, x2)
	return x
}

func isMax(x []int) bool {
	cnt := 0
	for _, v := range x {
		if v == 1 || v == 2 || v == 3 || v == 4 {
			cnt++
		}
	}
	var ret bool
	if cnt == len(x) {
		ret = false
	} else {
		ret = true
	}
	return ret
}

func main() {
	x := []int{0}
	fmt.Scanf("%d", &x[0])
	for isMax(x) {
		for _, v := range x {
			if v == 1 || v == 2 || v == 3 || v == 4 {
				// do nothing
			} else {
				x = generateX(x, v)
			}
		}
	}

	ret := 1
	for _, v := range x {
		ret = ret * v
	}
	if ret > 998244353 {
		fmt.Println(ret % 998244353)
	} else {
		fmt.Println(ret)
	}
	return
}
