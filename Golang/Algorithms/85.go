package Algorithms

import "fmt"

func Test85() {
	data1 := [][]byte{{'1', '0', '1', '0', '1'}, {'1', '0', '1', '1', '1'}, {'1','1','1','1','1'}, {'1','0','0','1','0'}}
	fmt.Println(maximalRectangle(data1), 6)
}


func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right, height := make([]int, cols), make([]int, cols), make([]int, cols)
	max := 0
	for i := 0; i < cols; i++ {
		right[i] = cols
	}
	for i := 0; i < rows; i++ {
		curLeft, curRight := 0, cols
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				height[j]++
				if left[j] < curLeft {
					left[j] = curLeft
				}
			} else {
				height[j] = 0
				left[j] = 0
				curLeft = j + 1
			}
		}
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if right[j] > curRight {
					right[j] = curRight
				}
			} else {
				right[j] = cols
				curRight = j
			}
		}
		for j := 0; j < cols; j ++ {
			area := (right[j] - left[j]) * height[j]
			if area > max {
				max = area
			}
		}
	}
	return max
}


func maximalRectangle2(matrix [][]byte) int {
	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				nextI := i
				for nextI < len(matrix) {
					nextJ := j + 1
					for nextJ  < len(matrix[i]) {
						amdYes := true
						for k := i; k <= nextI; k++ {
							if matrix[k][nextJ] != '1' {
								amdYes = false
								break
							}
						}
						if amdYes == false {
							break
						}
						nextJ++
					}
					area := (nextJ - j) * (nextI - i + 1)
					// fmt.Println( i, j,nextI, nextJ, area)
					if area > max {
						max = area
					}
					if nextI + 1 < len(matrix) &&
						matrix[nextI + 1][j] == '1' {
						nextI++
					} else {
						break
					}
				}
			}
		}
	}
	return max
}
