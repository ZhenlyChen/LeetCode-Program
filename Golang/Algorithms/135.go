package Algorithms

import "fmt"

func Test135() {
	fmt.Println(candy([]int{1, 0, 2}), 5)
	fmt.Println(candy([]int{1, 2, 2}), 4)
	fmt.Println(candy([]int{1,3,2,2,1}), 7)
	fmt.Println(candy([]int{1}), 1)
	fmt.Println(candy([]int{}), 0)
	fmt.Println(candy([]int{1,6,10,8,7,3,2}), 18)
}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	left, right := []int{1},[]int{1}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i - 1] {
			left = append(left, left[len(left) - 1] + 1)
		} else {
			left = append(left, 1)
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			right = append(right, right[len(right) - 1] + 1)
		} else {
			right = append(right, 1)
		}
	}
	sum := 0
	for i := range left {
		if left[i] > right[len(right) - i - 1] {
			sum += left[i]
		} else {
			sum += right[len(right) - i - 1]
		}
	}
	return sum
}