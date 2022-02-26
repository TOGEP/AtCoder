package main

import "fmt"

func searchSlice(slice []int, key int) []int {
	for i, v := range slice {
		if v == key {
			slice = append(slice[:i], slice[i+1:]...)
			return slice
		}
	}
	return nil
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	for i := 0; i < m; i++ {
		a = searchSlice(a, b[i])
		if a == nil {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}
