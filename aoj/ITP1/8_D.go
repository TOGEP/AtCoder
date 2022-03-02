package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var p string
	fmt.Scan(&p)

	// strings.Containsは第1引数の中に第2引数の文字列が含まれているかチェックする関数
	// 今回は時計回りに連続した文字列sを想定しているのでs+sを第1引数にすることで抜けなく検索できる
	// https://qiita.com/sayama0402/items/e32814e38375fafa919a
	if strings.Contains(s+s, p) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
