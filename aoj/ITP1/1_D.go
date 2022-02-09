package main

import "fmt"

func parseHour(s int) int {
	return s / 3600
}

func parseMinute(s int) int {
	return s % 3600 / 60
}

func parseSecond(s int) int {
	return s % 60
}

func main() {
	var s int
	fmt.Scanf("%d", &s)

	// \nをつけないとPEになるので注意
	fmt.Printf("%d:%d:%d\n", parseHour(s), parseMinute(s), parseSecond(s))
	return
}
