package main

func main() {

}

func generateMatrix(n int) [][]int {
	rowStart, rowEnd := 0, n-1
	colStart, colEnd := 0, n-1

	totalCount := n * n
	num := 1

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for num <= totalCount {
		if rowStart <= rowEnd {
			for j := colStart; j <= colEnd; j++ {
				res[rowStart][j] = num
				num++
			}
			rowStart++
		}

		if colStart <= colEnd {
			for i := rowStart; i <= rowEnd; i++ {
				res[i][colEnd] = num
				num++
			}
			colEnd--
		}

		if rowStart <= rowEnd {
			for j := colEnd; j >= colStart; j-- {
				res[rowEnd][j] = num
				num++
			}
			rowEnd--
		}

		if colStart <= colEnd {
			for i := rowEnd; i >= rowStart; i-- {
				res[i][colStart] = num
				num++
			}
			colStart++
		}
	}

	return res
}
