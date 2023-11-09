package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/dp/tasks/dp_c
func main() {
	sc := bufio.NewScanner(os.Stdin)

	inputN := readString(sc)
	n, _ := strconv.Atoi(inputN)

	for i := 0; i < n; i++ {
		inputABC := readStringAsSlice(sc)
		a, _ := strconv.Atoi(inputABC[0])
		b, _ := strconv.Atoi(inputABC[1])
		c, _ := strconv.Atoi(inputABC[2])

	}

	fmt.Println(ans)
}

// general util funcs
func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
func readStringAsSlice(sc *bufio.Scanner) []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}
