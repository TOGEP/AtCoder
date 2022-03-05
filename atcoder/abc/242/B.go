package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scan(&s)

	arr := make([]string, 0)
	for _, v := range s {
		arr = append(arr, string(v))
	}
	sort.Strings(arr)
	for _, v := range arr {
		fmt.Printf("%s", string(v))
	}
	return
}
