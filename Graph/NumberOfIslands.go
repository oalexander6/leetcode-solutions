package main

type Position struct {
	Row int
	Col int
}

func numIslands(grid [][]byte) int {
	numIslands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if string(grid[i][j]) == "0" {
				continue
			}

			numIslands++

			markIslandVisitedBFS(grid, i, j)
		}
	}

	return numIslands
}

func markIslandVisitedBFS(grid [][]byte, startRow int, startCol int) {
	stack := make([]Position, 1)
	stack[0] = Position{
		Row: startRow,
		Col: startCol,
	}

	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		grid[curr.Row][curr.Col] = []byte("0")[0]

		positionsToCheck := []Position{
			{Row: curr.Row + 1, Col: curr.Col},
			{Row: curr.Row - 1, Col: curr.Col},
			{Row: curr.Row, Col: curr.Col + 1},
			{Row: curr.Row, Col: curr.Col - 1},
		}

		for _, position := range positionsToCheck {
			if isValidCoordinate(position.Row, position.Col, len(grid), len(grid[0])) {
				if string(grid[position.Row][position.Col]) == "1" {
					grid[position.Row][position.Col] = []byte("0")[0]
					stack = append(stack, position)
				}
			}
		}
	}
}

func markIslandVisitedDFS(grid [][]byte, row int, col int) {
	grid[row][col] = []byte("0")[0]

	positionsToCheck := []Position{
		{Row: row + 1, Col: col},
		{Row: row - 1, Col: col},
		{Row: row, Col: col + 1},
		{Row: row, Col: col - 1},
	}

	for _, position := range positionsToCheck {
		if isValidCoordinate(position.Row, position.Col, len(grid), len(grid[0])) {
			if string(grid[position.Row][position.Col]) == "1" {
				grid[position.Row][position.Col] = []byte("0")[0]
				markIslandVisitedDFS(grid, position.Row, position.Col)
			}
		}
	}
}

func isValidCoordinate(row, col, numRows, numCols int) bool {
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}

// func main() {
// 	input := [][]byte{
// 		{[]byte("1")[0], []byte("1")[0], []byte("1")[0], []byte("1")[0], []byte("0")[0]},
// 		{[]byte("1")[0], []byte("1")[0], []byte("0")[0], []byte("1")[0], []byte("0")[0]},
// 		{[]byte("1")[0], []byte("1")[0], []byte("0")[0], []byte("0")[0], []byte("0")[0]},
// 		{[]byte("0")[0], []byte("0")[0], []byte("0")[0], []byte("0")[0], []byte("0")[0]},
// 	}

// 	_ = numIslands(input)
// }
