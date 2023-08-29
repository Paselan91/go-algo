package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc051/tasks/abc051_b
func main() {
	sc := bufio.NewScanner(os.Stdin)

	inputKS := readStringAsSlice(sc)
	k, _ := strconv.Atoi(inputKS[0])
	s, _ := strconv.Atoi(inputKS[1])

	fmt.Println(countEqualS(s, k))
}

func countEqualS(s, k int) int {
	count := 0
	for x := 0; x <= k; x++ {
		if x > s {
			break
		}
		for y := 0; y <= k; y++ {
			if x+y > s {
				break
			}

			z := s - x - y
			if z > k || z < 0 {
				continue
			}

			if x+y+z == s {
				count++
			}
		}
	}
	return count
}

// general util funcs
// func readString(sc *bufio.Scanner) string {
// 	sc.Scan()
// 	return sc.Text()
// }
func readStringAsSlice(sc *bufio.Scanner) []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}
