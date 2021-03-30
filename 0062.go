package main

import "fmt"

// dynamic programming programming
func uniquePaths(m int, n int) int {
	// initialize n x m matrix
	dp := make([][]int, n)
	for i:=0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	// fill row 0 and column with 1, since they have only one way to move
	for i:=0; i < m; i++ { // for row 0, the robot can only go down
		dp[0][i] = 1
	}
	for i:=0; i < n; i++ { // for column 0, the robot can only go right
		dp[i][0] = 1
	}

	for i:=1; i < n; i++ {
		for j:=1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]
}

func main(){
	m,n:=3,2
	dp := make([][]int, n)
	for i:=0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	fmt.Println(dp) // [[0 0 0] [0 0 0]]

	for i:=0; i < m; i++ { // for row 0, the robot can only go down
		dp[0][i] = 1
	}
	fmt.Println(dp) // [[1 1 1] [0 0 0]]

	for i:=0; i < n; i++ {	// for column 0, the robot can only go right
		dp[i][0] = 1
	}
	fmt.Println(dp) // [[1 1 1] [1 0 0]]
}