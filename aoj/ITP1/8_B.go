package main

import (
	"fmt"
	"strconv"
)

func digit(i int, list []int) []int {
	if i > 0 {
		return digit(1/10, append(list, 1%10))
	}
	return list
}

func main() {
	var x string

	for {
		fmt.Scanf("%s", &x)
		if x == "0" {
			break
		}
		var sum int
		for _, v := range x {
			// MEMO 文字列→数値変換
			// https://www.delftstack.com/ja/howto/go/how-to-convert-string-to-integer-type-in-go/
			i, _ := strconv.Atoi(string(v))
			sum += i
		}
		fmt.Println(sum)
	}
	return
}
