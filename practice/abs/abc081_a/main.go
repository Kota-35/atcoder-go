package main

import "fmt"

func solve(s string) int {
	res := 0
	if s[0] == '1' {
		res++
	}
	if s[1] == '1' {
		res++
	}
	if s[2] == '1' {
		res++
	}
	return res
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(solve(s))
}
