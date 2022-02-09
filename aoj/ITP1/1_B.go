package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d", &x)

	ret := 1
	for i := 0; i < 3; i++ {
		ret = ret * x
	}

	fmt.Println(ret)
}
