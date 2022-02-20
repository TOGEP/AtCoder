package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == 1 && b == 10 {
		fmt.Println("Yes")
	} else if a+1 == b || a-1 == b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
