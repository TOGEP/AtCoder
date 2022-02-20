package main

import "fmt"

func sliceUnique(a []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range a {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	s := sliceUnique(a)
	fmt.Println(len(s))
	return
}
