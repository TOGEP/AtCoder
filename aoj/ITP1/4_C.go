package main

import "fmt"

func main() {
	var a, b int
	var op string
	var ret int

	// Labeled Break
Calculation:
	for {
		fmt.Scanf("%d %s %d", &a, &op, &b)
		switch op {
		case "+":
			ret = a + b
		case "-":
			ret = a - b
		case "*":
			ret = a * b
		case "/":
			ret = a / b
		default:
			break Calculation
		}
		fmt.Println(ret)
	}
	return
}
