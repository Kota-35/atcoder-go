// https://atcoder.jp/contests/abs/tasks/abc087_b¥
package main

import "fmt"

func solve(a int, b int, c int, x int) int {
	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			rem := x - 500*i - 100*j
			if rem < 0 {
				continue
			}
			if rem%50 != 0 { // 50で割り切れない
				continue
			}
			k := rem / 50
			if k <= c { // 持っている枚数以内
				count++
			}
		}
	}

	return count
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	fmt.Println(solve(a, b, c, x))
}
