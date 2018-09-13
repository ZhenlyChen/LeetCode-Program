package Algorithms

import (
	"fmt"
	"math"
)

func Test4() {
	ans1 := findMedianSortedArrays([]int{1, 3}, []int{2})
	ans2 := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	ans3 := findMedianSortedArrays([]int{}, []int{1})
	ans4 := findMedianSortedArrays([]int{2}, []int{})
	ans5 := findMedianSortedArrays([]int{1, 5}, []int{2, 3, 4, 6})
	ans6 := findMedianSortedArrays([]int{}, []int{2, 3})
	fmt.Println(ans1, ans2, ans3, ans4, ans5, ans6)
	fmt.Println(2, 2.5, 1, 2, 3.5, 2.5)
}

func findMedianSortedArraysOld(nums1 []int, nums2 []int) float64 {
	count := len(nums1) + len(nums2)
	one, two, i, j := 0, 0, 0, 0
	for k := 0; k <= count/2; k++ {
		if j >= len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			if k == count/2-1 {
				one = nums1[i]
			} else if k == count/2 {
				two = nums1[i]
			}
			i++
		} else {
			if k == count/2-1 {
				one = nums2[j]
			} else if k == count/2 {
				two = nums2[j]
			}
			j++
		}
	}
	if count%2 == 0 {
		return float64(one+two) / 2
	}
	return float64(two)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	// make sure n > m
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m + n + 1) / 2 - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			left, right := 0, 0
			if i == 0 {
				left = nums2[j- 1]
			} else if j == 0 {
				left = nums1[i - 1]
			} else {
				left = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}
			if (m+n)%2 == 1 {
				return float64(left)
			}
			if i == m {
				right = nums2[j]
			} else if j == n {
				right = nums1[i]
			} else {
				right =int(math.Min(float64(nums2[j]), float64(nums1[i])))
			}
			return float64(left + right) / 2.0
		}
	}
	return 0.0
}
