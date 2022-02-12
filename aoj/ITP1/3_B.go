package main

import "fmt"

func main() {
	var x int
	for i := 0; ; i++ {
		fmt.Scanf("%d", &x)
		if x == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", i+1, x)
	}
	return
}
