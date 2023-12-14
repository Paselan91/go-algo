package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/generate-parentheses/description/

/*
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]
*/
func main() {
	n := 3
	// n := 1

	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	var res []string

	helper(n, 0, 0, "", &res)

	return res
}
func check(s string) bool {
	for len(s) > 0 {
		originalLength := len(s)
		s = strings.Replace(s, "()", "", 1)
		if len(s) == originalLength {
			return false
		}
	}
	return len(s) == 0
}

func helper(n int, openCount int, closeCount int, currentStr string, result *[]string) {
	if openCount == n && closeCount == n {
		if check(currentStr) {
			*result = append(*result, currentStr)
		}
		return
	}

	if openCount < n {
		helper(n, openCount+1, closeCount, currentStr+"(", result)
	}

	if closeCount < n {
		helper(n, openCount, closeCount+1, currentStr+")", result)
	}
}
