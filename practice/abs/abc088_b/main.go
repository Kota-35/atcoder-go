// https://atcoder.jp/contests/abs/tasks/abc088_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(as []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	alice, bob := 0, 0
	for i, a := range as {
		if i%2 == 0 {
			alice += a
		} else {
			bob += a
		}
	}
	return alice - bob
}

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	var as []int
	for _, input := range inputs {
		a, _ := strconv.Atoi(input)
		as = append(as, a)
	}

	fmt.Println(solve(as))
}
