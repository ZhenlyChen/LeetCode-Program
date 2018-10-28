package Algorithms

import (
	"fmt"
)

func Test57() {
	fmt.Println(insert([]Interval{{1,3}, {6,9}}, Interval{2, 5}))
	fmt.Println(insert([]Interval{{1,2}, {3,5}, {6,7}, {8,10}, {12,16}}, Interval{4, 8}))
	fmt.Println(insert([]Interval{}, Interval{4, 8}))
}

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	var res []Interval
	var i int
	for i < len(intervals) && intervals[i].End < newInterval.Start {
		i++
	}
	res = append(res, intervals[:i]...)
	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		if intervals[i].Start  < newInterval.Start {
			newInterval.Start = intervals[i].Start
		}
		if intervals[i].End > newInterval.End {
			newInterval.End = intervals[i].End
		}
		i++
	}
	res = append(res, newInterval)
	res = append(res, intervals[i:]...)
	return res
}

func insert2(intervals []Interval, newInterval Interval) []Interval {
	var res []Interval
	if len(intervals) == 0 {
		return res
	}
	var start, end int
	isAdd := false
	intervals = append(intervals, intervals[len(intervals) - 1])
	for _, c := range intervals {
		if end < min(newInterval.Start, c.Start) {
			if end - start > 0 {
				res = append(res, Interval{start, end})
			}
			if newInterval.Start > c.Start || isAdd {
				start = c.Start
				end = c.End
			} else {
				start = newInterval.Start
				end = newInterval.End
				isAdd = true
			}
		} else if end >= newInterval.Start && !isAdd {
			end = newInterval.End
			isAdd = true
		} else if end >= c.Start {
			end = c.End
		}
		fmt.Println(c, newInterval, start, end, res)
	}
	if end - start > 0 {
		res = append(res, Interval{start, end})
	}
	return res
}


func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}


func max(a,b int) int {
	if a < b {
		return b
	}
	return a
}