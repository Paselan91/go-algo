package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
func main() {
	startTime := []int{1, 2, 3, 4}
	endTime := []int{3, 2, 7, 5}
	queryTime := 4
	fmt.Println(busyStudent(startTime, endTime, queryTime))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	counter := 0
	for i, sTime := range startTime {
		if queryTime >= sTime && queryTime <= endTime[i] {
			counter++
		}
	}
	return counter
}
