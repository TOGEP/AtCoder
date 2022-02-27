package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func query1(a []int, x int) []int {
	a = append(a, x)
	return a
}

func query2(a []int, x int, k int) {
	var slice []int
	for _, v := range a {
		if v <= x {
			slice = append(slice, v)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	if len(slice) < k {
		fmt.Println(-1)
	} else {
		fmt.Println(slice[k-1])
	}
}

func query3(a []int, x int, k int) {
	var slice []int
	for _, v := range a {
		if v >= x {
			slice = append(slice, v)
		}
	}

	sort.Ints(slice)
	if len(slice) < k {
		fmt.Println(-1)
	} else {
		fmt.Println(slice[k-1])
	}
	return
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var a []int
	var q int
	fmt.Scan(&q)

	var input int

	sc.Split(bufio.ScanWords)
	for i := 0; i < q; i++ {
		input = nextInt()
		switch input {
		case 1:
			x := nextInt()
			a = query1(a, x)
		case 2:
			x := nextInt()
			k := nextInt()
			query2(a, x, k)
		case 3:
			x := nextInt()
			k := nextInt()
			query3(a, x, k)
		}
	}
	return
}
