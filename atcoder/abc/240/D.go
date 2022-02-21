package main

import "fmt"

type Stack struct {
	data []Pair
	size int
}

type Pair struct {
	a int
	b int
}

// Pushは新たにxを追加
func (s *Stack) Push(n Pair) {
	s.data = append(s.data, n)
	s.size++
}

// Popは最後に追加されたものを取り除く
func (s *Stack) Pop() bool {
	if s.size == 0 {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

// Topは最後に追加されたものを取得
func (s *Stack) Top() Pair {
	return s.data[s.size-1]
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var s Stack
	var cnt int
	for i := 0; i < n; i++ {
		cnt++
		if s.size == 0 || s.Top().a != a[i] {
			s.Push(Pair{a: a[i], b: 1})
		} else {
			k := s.Top()
			k.b += 1
			s.Pop()
			s.Push(k)
			if s.Top().b == a[i] {
				cnt -= s.Top().b
				s.Pop()
			}
		}
		fmt.Println(cnt)
	}
	return
}
