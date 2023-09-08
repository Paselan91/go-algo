package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func main() {
	ans := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(ans)
}
func lengthOfLongestSubstring(s string) int {
	longest := ""
	for _, char := range s {
		str := string(char)

		if strings.Contains(longest, str) {
			continue
		} else {
			
		}
		longest += str
	}

	fmt.Println(longest)
	return len(longest)
}
