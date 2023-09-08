package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.com/problems/add-two-numbers/
func main() {
	// node13 := ListNode{
	// 	Val:  3,
	// 	Next: nil,
	// }
	// node12 := ListNode{
	// 	Val:  4,
	// 	Next: &node13,
	// }
	// node11 := ListNode{
	// 	Val:  2,
	// 	Next: &node12,
	// }

	// node23 := ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }
	// node22 := ListNode{
	// 	Val:  6,
	// 	Next: &node23,
	// }
	// node21 := ListNode{
	// 	Val:  5,
	// 	Next: &node22,
	// }

	node11 := ListNode{
		Val:  0,
		Next: nil,
	}
	node21 := ListNode{
		Val:  0,
		Next: nil,
	}
	ans := addTwoNumbers(&node11, &node21)
	fmt.Println(ans.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numberL1 := makeNumbersFromList(l1)
	numberL2 := makeNumbersFromList(l2)

	addedNumber := convertIntToSlice(numberL1 + numberL2)

	lastNode := &ListNode{}
	var prevNode *ListNode = nil
	for i := 0; i < len(addedNumber); i++ {
		if i < len(addedNumber)-1 {
			tmpNode := &ListNode{
				Val:  addedNumber[i],
				Next: prevNode,
			}
			prevNode = tmpNode
		} else {
			lastNode = &ListNode{
				Val:  addedNumber[i],
				Next: prevNode,
			}
		}
	}

	return lastNode
}

func makeNumbersFromList(node *ListNode) int {
	listSlice := convertToSlice(node)
	number := 0
	// make numbers by listSlice's reverse order
	for i := 0; i < len(listSlice); i++ {
		number += listSlice[i] * int(math.Pow(10, float64(i)))
	}
	return number
}

func convertToSlice(l1 *ListNode) []int {
	slice := []int{}
	for i := l1; i != nil; i = i.Next {
		slice = append(slice, i.Val)
	}
	return slice
}

func convertIntToSlice(num int) []int {
	numberStr := strconv.Itoa(num)
	digits := make([]int, len(numberStr))
	for i, char := range numberStr {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
	}

	return digits
}

