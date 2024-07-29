package main

type Entry struct {
	ReachesAtl    bool
	ReachesPac    bool
	VisitedForAtl bool
	VisitedForPac bool
}

// type Position struct {
//     Row int
//     Col int
// }

func pacificAtlantic(heights [][]int) [][]int {
	memo := createMemo(len(heights), len(heights[0]))

	for row := 1; row < len(heights); row++ {
		for col := 1; col < len(heights[0]); col++ {

		}
	}

	results := make([][]int, 0)
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[0]); col++ {
			if memo[row][col].ReachesAtl && memo[row][col].ReachesPac {
				results = append(results, []int{row, col})
			}
		}
	}
	return results
}

func dfs(memo [][]Entry, row int, col int) {

}

func createMemo(numRows, numCols int) [][]Entry {
	memo := make([][]Entry, numRows)

	for i := 0; i < numRows; i++ {
		memo[i] = make([]Entry, numCols)
		for j := 0; j < numCols; j++ {
			entry := Entry{}

			if i == 0 || j == 0 {
				entry.ReachesPac = true
				entry.VisitedForPac = true
			}

			if i == numRows-1 || j == numCols-1 {
				entry.ReachesAtl = true
				entry.VisitedForAtl = true
			}

			memo[i][j] = entry
		}
	}

	return memo
}

func main() {
	pacificAtlantic([][]int{{1}})
}
