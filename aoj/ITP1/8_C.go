package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	alphabet := make(map[rune]int)

	// 複数行にまたがる標準入力
	for sc.Scan() {
		str := strings.ToLower(sc.Text())
		for _, v := range str {
			alphabet[v]++
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c : %d\n", i, alphabet[i])
	}
}
