package main

import "fmt"

// ソート並列化用チャネル
func sortChan(list []int, ch chan []int) {
	ret := quickSort(list)
	ch <- ret
}

// クイックソート
// 降順に並び替え
func quickSort(num []int) []int {
	if len(num) < 2 {
		return num
	}

	pivot := num[0]

	left := []int{}
	right := []int{}

	for _, n := range num[1:] {
		if n > pivot {
			right = append(right, n)
		} else {
			left = append(left, n)
		}
	}

	rightchan := make(chan []int)
	leftchan := make(chan []int)
	go sortChan(right, rightchan)
	go sortChan(left, leftchan)

	rl := <-rightchan
	ll := <-leftchan

	return append(append(rl, pivot), ll...)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d[i])
	}

	sorted := quickSort(d)
	var cnt = 1
	for i := 0; i < n-1; i++ {
		if sorted[i] > sorted[i+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
