package main

func setZeroes(matrix [][]int) {
	colZeroFlag := 1

	for rowI := 0; rowI < len(matrix); rowI++ {
		for colI := 0; colI < len(matrix[0]); colI++ {
			if matrix[rowI][colI] == 0 {
				if rowI == 0 {
					colZeroFlag = 0
				} else {
					matrix[rowI][0] = 0
					matrix[0][colI] = 0
				}
			}
		}
	}

	for rowI := 1; rowI < len(matrix); rowI++ {
		for colI := 1; colI < len(matrix[0]); colI++ {
			if matrix[rowI][0] == 0 || matrix[0][colI] == 0 {
				matrix[rowI][colI] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if colZeroFlag == 0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

// func main() {
// 	input := [][]int{
// 		[]int{1, 2, 3, 4},
// 		[]int{5, 0, 7, 8},
// 		[]int{0, 10, 11, 12},
// 		[]int{13, 14, 15, 0},
// 	}

// 	setZeroes(input)
// }
