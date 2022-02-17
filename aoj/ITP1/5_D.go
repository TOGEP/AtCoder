package main

import "fmt"

func call(n int) {
	for i := 1; i <= n; i++ {
		digit := 0
		for num := i; num != 0; num /= 10 {
			digit++
		}
		if i%3 == 0 {
			fmt.Printf(" %d", i)
		} else if i%10 == 3 {
			fmt.Printf(" %d", i)
		} else {
			x := i
			for j := 0; j < digit; j++ {
				x = x / 10
				if x%10 == 3 {
					fmt.Printf(" %d", i)
					break
				}
			}
		}
	}
	fmt.Printf("\n")
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	call(n)
	return
}
