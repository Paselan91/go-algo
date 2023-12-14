package main

import (
	"fmt"
	"sort"

	"mymodule/leetcode/2335/other_solutions"
)

// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/

/*
Input: amount = [1,4,2]
Output: 4

Input: amount = [5,4,4]
Output: 7

Input: amount = [5,0,0]
Output: 5
*/
func main() {
	amount := []int{1, 4, 2}
	// amount := []int{5, 4, 4}
	// amount := []int{5,0,0}

	fmt.Println(other_solutions.FillCupsByMaxHeap(amount))
}

func FillCups(arr []int) int {
	sortArr := func() {
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	}

	out := 0

	sortArr()
	for arr[0] != 0 {
		arr[0]--
		if arr[1] != 0 {
			arr[1]--
		}

		out++
		sortArr()
	}

	return out
}
