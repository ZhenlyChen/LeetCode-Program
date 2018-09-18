package Algorithms

import (
	"fmt"
)

func Test42() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}), 6)
	fmt.Println(trap([]int{2,1,0,3,0,1,2}), 6)
}

func trap(height []int) int {
	all := len(height)
	lMax, rMax := 0, 0
	l, r := 0, all - 1
	water := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				water += lMax - height[l]
			}
			l++
		} else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				water += rMax - height[r]
			}
			r--
		}
	}
	return water
}

func trap2(height []int) int {
	all := len(height)
	if all < 3 {
		return 0
	}
	water := 0
	for i := 1; i < all - 1; i++ {
		l, r := i - 1, i + 1
		for j := l; j >= 0; j-- {
			if height[j] > height[l] {
				l = j
			}
		}
		for j := r; j < all; j++ {
			if height[j] > height[r] {
				r = j
			}
		}
		if height[l] > height[i] && height[i] < height[r] {
			if height[l] > height[r] {
				water += height[r] - height[i]
			} else {
				water += height[l] - height[i]
			}
		}
		// fmt.Println(i,":",  height[l], height[i], height[r],"->", water)
	}
	return water
}