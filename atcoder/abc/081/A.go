package main

import "fmt"

var N = 3

func main() {
	var s int

	//%bで2進数
	fmt.Scanf("%b", &s)

	cnt := 0
	for i := 0; i < N; i++ {
		if (s>>i)&1 == 1 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
