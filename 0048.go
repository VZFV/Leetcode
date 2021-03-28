package main

func rotate(matrix [][]int)  {
	row := len(matrix)
	limit := (row-1)/2

	for i:=0; i<= limit; i++ {
		for j:=i;j<row-1-i;j++ {
			temp:= matrix[i][j]
			matrix[i][j] = matrix[row-1-j][i]
			matrix[row-1-j][i] = matrix[row-1-i][row-1-j]
			matrix[row-1-i][row-1-j] = matrix[j][row-1-i]
			matrix[j][row-1-i] = temp
		}
	}
}
