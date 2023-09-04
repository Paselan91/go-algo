package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc114/tasks/abc114_c
func main() {
	sc := bufio.NewScanner(os.Stdin)

	n := readString(sc)

	fmt.Println(calc753(strToInt(n)))
}

func calc753(n int) int {
	counter := 0
	minNum := 356

	if n <= minNum {
		return counter
	}

	count753(n, 0, 0b000, &counter)

	return counter
}

func count753(n, num, use int, counter *int) {
	// base case
	if num > n {
		return
	}

	if use == 0b111 {
		*counter++
	}

	// 7
	count753(n, 10*num+7, use|0b001, counter)

	// 5
	count753(n, 10*num+5, use|0b010, counter)

	// 3
	count753(n, 10*num+3, use|0b100, counter)
}

// general util funcs
func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

// func readStringAsSlice(sc *bufio.Scanner) []string {
// 	sc.Scan()
// 	return strings.Split(sc.Text(), " ")
// }
