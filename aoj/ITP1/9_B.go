package main

import "fmt"

func main() {
	for {
		var str string
		fmt.Scan(&str)
		if str == "-" {
			break
		}
		queue := make([]string, 0)
		for _, v := range str {
			queue = append(queue, string(v))
		}

		var m int
		var h int
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scan(&h)
			//dequeue
			dq := queue[:h]
			queue = queue[h:]
			//enqueue
			queue = append(queue, dq...)
		}
		for _, v := range queue {
			fmt.Print(string(v))
		}
		fmt.Printf("\n")
	}
}
