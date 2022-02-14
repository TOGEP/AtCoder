package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	cnt := 0
	for i := a; i <= b; i++ {
		if c%i == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
