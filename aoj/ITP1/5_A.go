package main

import "fmt"

func main() {
	var h, w int

Loop:
	for {
		fmt.Scanf("%d %d", &h, &w)
		if h == 0 && w == 0 {
			break Loop
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	return
}
