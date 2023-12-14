package main

import (
	"fmt"
)

// https://leetcode.com/problems/contains-duplicate/

// Example 1:
// Input: nums = [1,2,3,1]
// Output: true

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false

// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{1,2,3,4}
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
	// fmt.Println(nums1[1:])
}

// sort ver
// func containsDuplicate(nums []int) bool {
// 	sort.Ints(nums)
// 	minNum := nums[0]
//     for _, num := range nums[1:] {
// 		if num == minNum{
// 			return true
// 		}
// 		minNum = num
// 	}
// 	return false
// }

// map ver
func containsDuplicate(nums []int) bool {
	numsCountMap := map[int]int{}
	for _, num := range nums {
		_, ok := numsCountMap[num]
		if !ok {
			numsCountMap[num] = 1
		} else {
			return true
		}
	}
	return false
}
