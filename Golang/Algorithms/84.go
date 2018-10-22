package Algorithms

import "fmt"

func Test84() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}), 10)
	fmt.Println(largestRectangleArea([]int{1, 1}), 2)
	fmt.Println(largestRectangleArea([]int{1}), 1)
	fmt.Println(largestRectangleArea([]int{}), 0)
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	max := 0
	stacks := make([]int, l + 1)
	is := -1
	for i := 0; i <= l; i++ {
		h := 0
		if i != l {
			h = heights[i]
		}
		for is != -1 && h < heights[stacks[is]] {
			hh := heights[stacks[is]]
			is--
			width := i
			if is != -1 {
				width = i - 1 - stacks[is]
			}
			if hh * width > max {
				max = hh * width
			}
		}
		is++
		stacks[is] = i
	}
	return max
}

func largestRectangleArea3(heights []int) int {
	l := len(heights)
	if l < 1 {
		return 0
	}
	leftLess := make([]int, l)
	rightLess := make([]int, l)
	leftLess[0] = -1
	rightLess[l - 1] = l
	for i := 1; i < l; i++ {
		left := i - 1
		for left >= 0 && heights[left] >= heights[i] {
			left = leftLess[left]
		}
		leftLess[i] = left
	}
	for i := l - 2; i >= 0; i-- {
		right := i + 1
		for right < l && heights[right] >= heights[i] {
			right = rightLess[right]
		}
		rightLess[i] = right
	}
	max := 0
	for i := 0; i < l; i++ {
		area := heights[i] * (rightLess[i] - leftLess[i] - 1)
		if area > max {
			max = area
		}
	}
	return max
}

func largestRectangleArea2(heights []int) int {
	max := 0
	for i := range heights {
		left, right := i, i
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			}
			right = j
		}
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			left = j
		}
		area := heights[i] * (right - left + 1)
		if area > max {
			max = area
		}
 	}
	return max
}