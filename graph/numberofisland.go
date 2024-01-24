package main

import "fmt"

var row = []int{0, 0, -1, 1}
var col = []int{-1, 1, 0, 0}
var mat = [][]byte{
	{'1', '1', '0', '0', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
}

func main() {
	fmt.Println(numberOfIsland(mat))
}
func numberOfIsland(mat [][]byte) int {
	count := 0
	r := len(mat)
	c := len(mat[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] != '0' {
				count = count + 1
				dfs(i, j, r, c, mat)
			}
		}
	}
	return count
}

func dfs(i, j, r, c int, mat [][]byte) {
	if i < 0 || i >= r || j < 0 || j >= c {
		return
	}
	if mat[i][j] == '0' {
		return
	}
	mat[i][j] = '0'
	for k := 0; k < 4; k++ {
		dfs(i+row[k], j+col[k], r, c, mat)
	}

}
