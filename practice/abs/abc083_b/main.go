// https://atcoder.jp/contests/abs/tasks/abc083_b
package main

import "fmt"

func calculateSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func solve(n int, a int, b int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if sum := calculateSum(i); sum >= a && sum <= b {
			res += i
		}
	}
	return res
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(solve(n, a, b))
}
