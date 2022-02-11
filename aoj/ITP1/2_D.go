package main

import "fmt"

func main() {
	var w, h, x, y, r int
	fmt.Scanf("%d %d %d %d %d", &w, &h, &x, &y, &r)

	// 円の上下左右の各頂点(x+r, y), (x-r, y), (x, y+r), (x, y-r)が(0,0)以上かつ(w,h)以下か
	if (x-r >= 0 && x+r <= w) && (y-r >= 0 && y+r <= h) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
