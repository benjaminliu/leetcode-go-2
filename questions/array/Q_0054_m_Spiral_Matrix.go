package main

import "github.com/benjaminliu/leetcode-go-2/utils"

func main() {
	matrics := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	utils.PrintList(spiralOrder(matrics))
}

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])

	leftBound := 0
	rightBound := col - 1
	upperBound := 0
	lowerBound := row - 1

	totalSize := row * col
	res := make([]int, 0)
	for len(res) < totalSize {
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upperBound][j])
			}
			upperBound++
		}

		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}

		if upperBound <= lowerBound {
			for j := rightBound; j >= leftBound; j-- {
				res = append(res, matrix[lowerBound][j])
			}
			lowerBound--
		}

		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}
	return res
}
