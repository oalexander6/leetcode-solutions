package main

func getRow(rowIndex int) []int {
    lastRow := make([]int, rowIndex+1)
    lastRow[0] = 1

    for row := 2; row <= rowIndex+1; row++ {
        for col := row-1; col >= 0; col-- {
            if col == 0 || col == row-1 {
                lastRow[col] = 1
                continue
            }

            lastRow[col] = lastRow[col] + lastRow[col-1]
        }
    }

    return lastRow
}
