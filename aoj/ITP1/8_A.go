package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	var str string
	str = sc.Text()

	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			fmt.Print(strings.ToUpper(string(v)))
		} else if v >= 'A' && v <= 'Z' {
			fmt.Print(strings.ToLower(string(v)))
		} else {
			fmt.Print(string(v))
		}
	}
	fmt.Printf("\n")
	return
}
