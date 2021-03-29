package main

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	rowBegin, rowEnd, columnBegin, columnEnd := 0, len(matrix)-1, 0, len(matrix[0])-1

	for rowBegin <= rowEnd && columnBegin <= columnEnd {
		for i:=columnBegin; i<=columnEnd;i++ {
			res = append(res, matrix[rowBegin][i])
		}
		rowBegin++
		for i:=rowBegin;i<=rowEnd;i++ {
			res = append(res, matrix[i][columnEnd]);
		}
		columnEnd--
		if rowBegin <= rowEnd {
			for i:=columnEnd; i >= columnBegin; i-- {
				res = append(res, matrix[rowEnd][i])
			}
		}
		rowEnd--;
		if columnBegin <= columnEnd {
			for i:=rowEnd; i>=rowBegin;i-- {
				res = append(res, matrix[i][columnBegin])
			}
		}
		columnBegin++
	}
	return res
}