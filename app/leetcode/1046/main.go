package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/last-stone-weight/description/
func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}

// sort ver
func lastStoneWeight(stones []int) int {
	isLeftOne := len(stones) == 1

	for !isLeftOne {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		x := stones[0]
		y := stones[1]
		stones = stones[2:]
		stones = append(stones, x-y)
		isLeftOne = len(stones) == 1
	}

	return stones[0]
}
