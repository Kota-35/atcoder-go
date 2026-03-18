// https://atcoder.jp/contests/abs/tasks/abc081_b

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func countTwos(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			count++
			num /= 2
		} else {
			break
		}
	}

	return count
}

func solve(as []int) int {
	res := math.MaxInt
	for _, a := range as {
		res = min(res, countTwos(a))
	}
	return res
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
