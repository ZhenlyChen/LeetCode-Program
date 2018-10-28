package Algorithms

import "fmt"

func Test51() {
	res := solveNQueens(4)
	for i := range res {
		for j := range res[i] {
			fmt.Println(res[i][j])
		}
		fmt.Println("---")
	}
	// fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	row := make([]int, n)
	col := make([]bool, n)
	left := make([]bool, 2*n)
	right := make([]bool, 2*n)
	res :=  make([][]string, 0)
	findNQueens(n,0 ,&row, &col, &left, &right, &res)
	return res
}

func findNQueens(n, r int, row *[]int, col, left, right *[]bool, res *[][]string) {
	if r == n {
		var solve []string
		blankRow := make([]byte, n)
		for j := 0; j < n; j++ {
			blankRow[j] = '.'
		}
		for i := range *row {
			newRow := make([]byte, n)
			copy(newRow, blankRow)
			newRow[(*row)[i]] = 'Q'
			solve = append(solve, string(newRow))
		}
		*res = append(*res, solve)
		return
	}
	for c := 0; c < n; c++ {
		if (*col)[c] == false && (*left)[n - r + c] == false && (*right)[c + r + 1] == false{
			(*row)[r] = c
			(*col)[c] = true
			(*left)[n - r + c] = true
			(*right)[c + r + 1] = true
			findNQueens(n, r + 1, row, col, left, right ,res)
			(*col)[c] = false
			(*left)[n - r + c] = false
			(*right)[c + r + 1] = false
		}
	}
}