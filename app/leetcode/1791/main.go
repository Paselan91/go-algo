package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-center-of-star-graph/

/*
Input: edges = [[1,2],[2,3],[4,2]]
Output: 2

Input: edges = [[1,2],[5,1],[1,3],[1,4]]
Output: 1
*/
func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	// amount := []int{5, 4, 4}
	// amount := []int{5,0,0}

	fmt.Println(findCenter(edges))
}

func findCenter(edges [][]int) int {
	var prevEdge1, prevEdge2 int

	hasValueInSlice := func(target int, values []int) bool {
		for _, v := range values {
			if v == target {
				return true
			}
		}
		return false
	}

	for i, edge := range edges {
		if i > 0 {
			prevEdge1 = edge[0]
			prevEdge2 = edge[1]
			continue
		}
		if !hasValueInSlice(prevEdge1, edge) {
			prevEdge1 = edge[0]
		}
		if hasValueInSlice(prevEdge2, edge) {
			prevEdge2 = edge[1]
		}
	}

	return prevEdge1
}
