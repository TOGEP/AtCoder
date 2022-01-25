package main

import (
	"fmt"
)

func swap(a *string, b *string) {
	tmp := *a
	*a = *b
	*b = tmp
}

// クイックソート
// Q = 1000は間に合うという前提 -> Qの値による終了処理は行わない
// TODO 毎回最小値がヒポットとして選ばれているのでO(n^2)...
//      N個のラベル付けをランダムに取得して平均O(n log n)にしたい
// TODO Q = 7の時
func quickSort(str []string, lenght int) (ret []string) {
	if lenght < 2 {
		return str
	}

	pivot := str[0]

	left := []string{}
	right := []string{}

	for _, v := range str[1:] {
		fmt.Printf("? %s %s\n", v, pivot)
		var ans string
		fmt.Scanf("%s", &ans)
		if ans == ">" {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	rightchan := make(chan []string)
	leftchan := make(chan []string)
	go sortChan(right, rightchan)
	go sortChan(left, leftchan)

	rl := <-rightchan
	ll := <-leftchan

	return append(append(ll, pivot), rl...)
}

func sortChan(list []string, ch chan []string) {
	ret := quickSort(list, len(list))
	ch <- ret
}

func main() {
	var N, Q int

	fmt.Scanf("%d %d", &N, &Q)

	str := make([]string, N)

	for i, s := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		if i == N {
			break
		}
		str[i] = string(s)
	}

	ret := quickSort(str, N)

	fmt.Printf("! ")
	for _, v := range ret {
		fmt.Printf("%s", v)
	}
	fmt.Printf("\n")
}
