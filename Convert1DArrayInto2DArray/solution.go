package main

func construct2DArray(original []int, m int, n int) [][]int {
	results := make([][]int, m)
	numNumbers := len(original)

	for i := 0; i < numNumbers; i++ {
		rowNum := (i / m)
		colNum := 0

		if i != 0 {
			colNum = i % n
		}

		if colNum == 0 {
			results[rowNum] = make([]int, n)
		}

		results[rowNum][colNum] = original[i]
	}
	return results
}

func main() {
	construct2DArray([]int{1, 2, 3}, 1, 3)
}
