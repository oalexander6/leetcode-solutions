package graph

type position struct {
	Row int
	Col int
}

func pacificAtlantic(heights [][]int) [][]int {
	NUM_ROWS := len(heights)
	NUM_COLS := len(heights[0])

	visitedPac := make(map[position]bool)
	visitedAtl := make(map[position]bool)

	for row := 0; row < NUM_ROWS; row++ {
		dfs(row, 0, visitedPac, heights[row][0], heights)
		dfs(row, NUM_COLS-1, visitedAtl, heights[row][NUM_COLS-1], heights)
	}

	for col := 0; col < NUM_COLS; col++ {
		dfs(0, col, visitedPac, heights[0][col], heights)
		dfs(NUM_ROWS-1, col, visitedAtl, heights[NUM_ROWS-1][col], heights)
	}

	results := make([][]int, 0)
	for pos, _ := range visitedPac {
		if _, isVisitedBoth := visitedAtl[pos]; isVisitedBoth {
			results = append(results, []int{pos.Row, pos.Col})
		}
	}
	return results
}

func dfs(row int, col int, visited map[position]bool, lastHeight int, heights [][]int) {
	currPos := position{
		Row: row,
		Col: col,
	}
	_, isVisited := visited[currPos]

	if isVisited || !isValidPosition(row, col, len(heights), len(heights[0])) || heights[row][col] < lastHeight {
		return
	}

	visited[currPos] = true

	neighbors := []position{
		{Row: row + 1, Col: col},
		{Row: row - 1, Col: col},
		{Row: row, Col: col + 1},
		{Row: row, Col: col - 1},
	}

	for _, pos := range neighbors {
		dfs(pos.Row, pos.Col, visited, heights[row][col], heights)
	}
}

func isValidPosition(row int, col int, numRows int, numCols int) bool {
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}
