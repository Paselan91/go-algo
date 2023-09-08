package main

import (
	"fmt"
)

// https://leetcode.com/problems/two-sum/
func main() {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
    // Space: O(n)
    m := make(map[int]int, 0)

    // Time: O(n)
    for idx, num := range nums {
        m[num] = idx
    }

    // Time: O(n)
    for indexX, num := range nums {
        // Time: O(1)
        if indexY, ok := m[target - num]; ok && indexX != indexY {
            return []int{indexX, indexY}
        }
    }

    return []int{}
}
