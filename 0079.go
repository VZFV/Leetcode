package main

func exist(board [][]byte, word string) bool {
	//visited := make([][]bool, len(board))		// record the visited index
	//for i:=0; i < len(visited); i++ {
	//	visited[i] = make([]bool, len(board[0]))
	//}
	rows, columns := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j:= 0; j < columns; j++ {
			if word[0] == board[i][j] && searchWorld(i, j, board, 0, word) {
				return true
			}
		}
	}
	return false
}

func searchWorld(i int, j int, board [][]byte, index int, word string) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || word[index] != board[i][j] {
		return false
	}
	temp := board[i][j]				// save the current index of board
	board[i][j] = 0					// set the current index value to 0 to prevent recursion from accessing again
	if searchWorld(i+1, j, board, index+1, word) == true ||
	   searchWorld(i-1, j, board, index+1, word) == true ||
	   searchWorld(i, j+1, board, index+1, word) == true ||
	   searchWorld(i, j-1, board, index+1, word) == true {
		return true
	}

	board[i][j] = temp				/// set back the actual value

	return false
}