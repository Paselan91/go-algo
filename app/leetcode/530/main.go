package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(getMinimumDifference(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	isLastNode := false
	var parent, left, right, minDiff float64
	for isLastNode {
		parent = float64(root.Val)
		if root.Left != nil {
			left = float64(root.Left.Val)
			tmpDiff := math.Abs(parent - left)
			if tmpDiff < minDiff {
				minDiff = tmpDiff
			}
		}
		if root.Right != nil {
			right = float64(root.Right.Val)
			tmpDiff := math.Abs(parent - right)
			if tmpDiff < minDiff {
				minDiff = tmpDiff
			}
		}
		root = root.Left
		if root.Left == nil {
			root = root.Right
		}
	}
	return int(minDiff)
}
