package main

import "fmt"

func main() {
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}

	var n1 int
	var n2 int
	var n3 int
	var n4 int
	for i, v := range a {
		if v == 0 {
			n1 = a[i]
		}
	}
	for i, v := range a {
		if v == a[n1] {
			n2 = a[i]
		}
	}
	for i, v := range a {
		if v == a[n2] {
			n3 = a[i]
		}
	}
	for i, v := range a {
		if v == a[n3] {
			n4 = a[i]
		}
	}
	fmt.Println(n4)
	return
}
