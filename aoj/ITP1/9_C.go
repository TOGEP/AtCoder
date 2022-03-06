package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	//score[0] is tarou score
	//score[1] is haanako score
	score := make([]int, 2)
	for i := 0; i < n; i++ {
		var a string
		fmt.Scan(&a)
		var b string
		fmt.Scan(&b)
		if a > b {
			score[0] += 3
		} else if a < b {
			score[1] += 3
		} else {
			score[0]++
			score[1]++
		}
	}
	fmt.Printf("%d %d\n", score[0], score[1])
	return
}
