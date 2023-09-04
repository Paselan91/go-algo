package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc045/tasks/arc061_a?lang=ja
func main() {
	sc := bufio.NewScanner(os.Stdin)

	input := readString(sc)

	fmt.Println(calcS(input))
}

func calcS(S string) int64 {
	N := len(S)

	var res int64 = 0
	for bit := 0; bit < (1 << (N - 1)); bit++ {
		var tmp int64 = 0
		for i := 0; i < N-1; i++ {
			tmp = tmp*10 + int64(stringToInt(string(S[i])))
			if (bit & (1 << i)) > 0 {
				res += tmp
				tmp = 0
				fmt.Println("in if")
			}
			fmt.Printf("i:%03b, tmp:%d, res:%d\n", i, tmp, res)
			fmt.Println("----------")
		}
		tmp = tmp*10 + int64(stringToInt(string(S[len(S)-1])))
		res += tmp
		fmt.Printf("tmp:%d\n", tmp)
		fmt.Printf("res:%d\n", res)
		fmt.Println("==================")
	}

	return res
}

// general util funcs
func stringToInt(s string) int {
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
