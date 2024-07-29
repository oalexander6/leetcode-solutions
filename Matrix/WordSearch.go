package main

type position struct {
	Row int
	Col int
}

func exist(board [][]byte, word string) bool {
	numRows := len(board)
	numCols := len(board[0])

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if board[row][col] == word[0] {
				currPos := position{Row: row, Col: col}
				visited := make(map[position]bool)
				visited[currPos] = true
				if doSearch(currPos, word[1:], board, visited) {
					return true
				}
			}
		}
	}

	return false
}

func doSearch(pos position, word string, board [][]byte, visited map[position]bool) bool {
	if len(word) == 0 {
		return true
	}

	candidates := []position{
		{Row: pos.Row + 1, Col: pos.Col},
		{Row: pos.Row - 1, Col: pos.Col},
		{Row: pos.Row, Col: pos.Col + 1},
		{Row: pos.Row, Col: pos.Col - 1},
	}

	for _, candidate := range candidates {
		if !isValidCoordinate(candidate, len(board), len(board[0])) {
			continue
		}
		if board[candidate.Row][candidate.Col] == word[0] {
			if _, isVisited := visited[candidate]; !isVisited {
				visited[candidate] = true
				if doSearch(candidate, word[1:], board, visited) {
					return true
				} else {
					delete(visited, candidate)
				}
			}
		}
	}
	return false
}

func isValidCoordinate(pos position, numRows, numCols int) bool {
	return pos.Row >= 0 && pos.Row < numRows && pos.Col >= 0 && pos.Col < numCols
}

func main() {
	input := [][]byte{
		{[]byte("A")[0], []byte("B")[0], []byte("C")[0], []byte("E")[0]},
		{[]byte("S")[0], []byte("F")[0], []byte("E")[0], []byte("S")[0]},
		{[]byte("A")[0], []byte("D")[0], []byte("E")[0], []byte("E")[0]},
	}

	exist(input, "ABCESEEEFS")
}
