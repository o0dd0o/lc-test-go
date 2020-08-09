package main

import "fmt"

func gameOfLife(board [][]int) {
	lens := len(board)
	lenl := len(board[0])
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < lens; i++ {
		for j := 0; j < lenl; j++ {
			sum := 0
			for k := 0; k < 8; k++ {
				x := dx[k] + j
				y := dy[k] + i

				if x >= 0 && x < lenl && y >= 0 && y < lens {
					sum += board[y][x] & 1
				}

			}

			if board[i][j] == 1 {
				if sum > 1 && sum < 4 {
					board[i][j] += 2
				}
			} else {
				if sum == 3 {
					board[i][j] += 2
				}
			}
		}
	}

	for i := 0; i < lens; i++ {
		for j := 0; j < lenl; j++ {
			board[i][j] = board[i][j] >> 1
		}
	}
}
func main() {
	b := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(b)
	fmt.Printf("%v", b)
}
