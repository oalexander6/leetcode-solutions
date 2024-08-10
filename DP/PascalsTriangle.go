package main

func generate(numRows int) [][]int {
    result := make([][]int, numRows)
    result[0] = []int{1}

    for row := 2; row <= numRows; row++ {
        rowData := make([]int, row)
        for col := 0; col < row; col++ {
            if col == 0 || col == row-1 {
                rowData[col] = 1
                continue
            }

            rowData[col] = result[row-2][col-1] + result[row-2][col]
        }
        result[row-1] = rowData
    }

    return result
}
