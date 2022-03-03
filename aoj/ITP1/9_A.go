package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)
	w = strings.ToLower(w)

	var sc = bufio.NewScanner(os.Stdin)

	var cnt int
	for sc.Scan() {
		str := strings.ToLower(sc.Text())
		arr := strings.Split(str, " ")
		for _, v := range arr {
			if v == w {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
