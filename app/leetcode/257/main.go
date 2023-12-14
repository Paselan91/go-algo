package main

import (
	"fmt"
)

// https://leetcode.com/problems/average-of-levels-in-binary-tree/

/*
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
*/
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

	fmt.Println(averageOfLevels(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelLength := len(queue)
		levelSum := 0

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levelAvg := float64(levelSum) / float64(levelLength)
		result = append(result, levelAvg)
	}

	return result
}
