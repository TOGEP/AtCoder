package main

import (
	"fmt"
	"regexp"
)

const N = 4

func main() {
	var s string
	fmt.Scanf("%s", &s)

	r := regexp.MustCompile(`^(dream|dreamer|erase|eraser)+$`)
	if r.MatchString(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	return
}
