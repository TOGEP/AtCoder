package main

import "fmt"

func main() {
	var n, x int
	for {
		fmt.Scan(&n, &x)
		if n == 0 && x == 0 {
			break
		}
		var cnt int
		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
	return
}
