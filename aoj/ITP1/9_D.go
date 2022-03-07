package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	str := []rune(s)

	var q int
	fmt.Scan(&q)

	var sc = bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		scan := sc.Text()
		input := strings.Split(scan, " ")
		a, _ := strconv.Atoi(input[1])
		b, _ := strconv.Atoi(input[2])
		switch input[0] {
		case "print":
			fmt.Println(string(str[a : b+1]))
		case "reverse":
			for i := 0; i <= (b-a)/2; i++ {
				str[a+i], str[b-i] = str[b-i], str[a+i]
			}

		case "replace":
			p := []rune(input[3])
			for i := 0; i <= b-a; i++ {
				str[a+i] = p[i]
			}
		}
	}
	return
}
