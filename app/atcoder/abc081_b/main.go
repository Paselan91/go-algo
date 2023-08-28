package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc081/tasks/abc081_b
func main() {
	sc := bufio.NewScanner(os.Stdin)

	inputN := readString(sc)
	n, _ := strconv.Atoi(inputN)
	numbers := readStringAsSlice(sc)

	ans := 0
	isAllEven := true
	for isAllEven {
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(numbers[i])
			if num%2 == 0 {
				numbers[i] = strconv.Itoa(num / 2)
			} else {
				isAllEven = false
				break
			}
		}
		if isAllEven {
			ans++
		}
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
