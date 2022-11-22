package main

import "github.com/benjaminliu/leetcode-go-2/utils"

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	numMatrix := Constructor(matrix)

	utils.Println(numMatrix.SumRegion(2, 1, 4, 3))
}

type NumMatrix struct {
	board [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	board := make([][]int, len(matrix)+1)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i < len(board); i++ {
		for j := 1; j < len(board[0]); j++ {
			board[i][j] = board[i][j-1] + matrix[i-1][j-1]
		}
	}

	for j := 1; j < len(board[0]); j++ {
		for i := 1; i < len(board); i++ {
			board[i][j] = board[i][j] + board[i-1][j]
		}
	}

	return NumMatrix{
		board: board,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.board[row2+1][col2+1] +
		this.board[row1][col1] -
		this.board[row2+1][col1] -
		this.board[row1][col2+1]
}
