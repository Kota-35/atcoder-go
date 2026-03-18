// https://atcoder.jp/contests/abs/tasks/abc085_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(ds []int) int {
	seen := map[int]bool{}
	for _, d := range ds {
		seen[d] = true
	}
	return len(seen)
}

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	var ds []int
	for i := 0; i < n; i++ {
		sc.Scan()
		d, _ := strconv.Atoi(sc.Text())
		ds = append(ds, d)
	}
	fmt.Println(solve(ds))
}
