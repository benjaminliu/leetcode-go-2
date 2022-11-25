package main

func main() {

}

func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			swap(matrix, i, j, j, i)
		}
	}

	lastCol := col - 1
	for i := 0; i < row; i++ {
		left, right := 0, lastCol
		for left < right {
			swap(matrix, i, left, i, right)
			left++
			right--
		}
	}
}

func swap(matrix [][]int, row1, col1, row2, col2 int) {
	tmp := matrix[row1][col1]
	matrix[row1][col1] = matrix[row2][col2]
	matrix[row2][col2] = tmp
}
