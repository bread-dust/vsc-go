package main

import "sort"

// www.leetcode-cn.com/problems/course-schedule-iii/
func scheduleCourse(course [][]int) int {
	sort.Slice(course, func(i, j int) bool {
		return course[i][1] < course[j][1]
	}

	h := &Heap{}
	time := 0

	for _, c := range course {
		if time+c[0] <= c[1] {
			time += c[0]
			h.Push(c[0])
		} else if h.Peek() > c[0] {
			time += c[0] - h.Pop()
			h.Push(c[0])
		}
	}
	return h.Len()
}
