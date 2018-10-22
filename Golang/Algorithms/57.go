package Algorithms

import (
	"fmt"
	"strconv"
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

func (i *Interval) ToString() string {
	return strconv.Itoa(i.Start) + "," + strconv.Itoa(i.End);
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	return make([]Interval, 1)
}
