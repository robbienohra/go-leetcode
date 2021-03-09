package main

func box(row, col int) int {

	return (row/3)*3 + col/3

}

func isValidSudoku(board [][]byte) bool {

	cols := make(map[int]map[byte]bool)
	rows := make(map[int]map[byte]bool)
	boxes := make(map[int]map[byte]bool)

	for i := 0; i < 9; i++ {
		cols[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {

		for col := 0; col < 9; col++ {

			item := board[row][col]

			if item == '.' {
				continue
			}

			index := box(row, col)

			box := boxes[index]
			row := rows[row]
			col := cols[col]

			if box[item] || row[item] || col[item] {
				return false
			}

			box[item] = true
			row[item] = true
			col[item] = true

		}
	}

	return true
}
