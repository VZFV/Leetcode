package main

import (
	"fmt"
	"sort"
)

// make can only be used with slice, map, channel
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1{
		return intervals
	}
	sort.Slice(intervals, func (i, j int) bool{
		return intervals[i][0] <= intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	currentIndex := 0

	for i:=1;i<len(intervals);i++{
		if res[currentIndex][1] >= intervals[i][0] {
			res[currentIndex][1] = max(res[currentIndex][1], intervals[i][1])
		} else {
			currentIndex++
			res = append(res, intervals[i])
		}
	}
	return res
}

func main(){
	temp := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(temp))

}

func max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}