package main

var visited [][]bool

func solveDfs(board [][]byte) {
	visited = make([][]bool, len(board))
	for i, row := range board {
		visited[i] = make([]bool, len(row))
	}

	for i, row := range board {
		for j, _ := range row {
			dfs(board, i, j)
		}
	}

	for i, row := range board {
		for j, _ := range row {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'E' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row int, col int) {
	if board[row][col] == 'O' &&
		row >= 0 && row < len(board) &&
		col >= 0 && col < len(board[row]) {
		board[row][col] = 'E'

		dfs(board, row+1, col+1)
		dfs(board, row+1, col-1)
		dfs(board, row-1, col+1)
		dfs(board, row-1, col-1)
	}
}
