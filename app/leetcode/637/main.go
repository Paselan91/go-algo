package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

/*
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
*/
func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(binaryTreePaths(tree))
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
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	preOrder(root, "", &res)
	return res
}

func preOrder(root *TreeNode, path string, res *[]string) {
	isLeaf := root.Left == nil && root.Right == nil
	if isLeaf {
		*res = append(*res, path+strconv.Itoa(root.Val))
	}

	path = path + strconv.Itoa(root.Val) + "->"
	if root.Left != nil {
		preOrder(root.Left, path, res)
	}
	if root.Right != nil {
		preOrder(root.Right, path, res)
	}
}
