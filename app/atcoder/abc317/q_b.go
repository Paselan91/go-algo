package main

/*
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	inputN := readString(sc)
	n, _ := strconv.Atoi(inputN)

	numbers := readStrings(sc)

	min := findMinFromStrSlice(numbers)

	ans := min
	for i := min; i < min + n + 1; i++ {
		// findしてnumbersをremoveすればもっと早くなる
		hasValue := findFromSlice(numbers, strconv.Itoa(i))
		if !hasValue {
			ans = i
		}
	}

	fmt.Println(ans)
}

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
func readStrings(sc *bufio.Scanner) []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}

func findMinFromStrSlice(stringSlice []string) int {
	minValue, _ := strconv.Atoi(stringSlice[0])

	for i, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("要素 %d: %s を数値に変換できませんでした。\n", i, str)
			continue
		}

		if i == 0 || num < minValue {
			minValue = num
		}
	}
	return minValue
}

func findFromSlice(stringSlice []string, value string) bool {
	for _, str := range stringSlice {
		if str == value {
			return true
		}
	}
	return false
}
*/