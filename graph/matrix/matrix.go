package matrix

import "fmt"

func Matrixrep(n int, edges [][]int) [6][6]int {
	var matrix [6][6]int
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = 1
		matrix[edge[1]][edge[0]] = 1
	}
	display(matrix)
	return matrix
}
func display(edges [6][6]int) {
	for _, edge := range edges {
		for _, val := range edge {
			fmt.Printf("%d ", val)
		}
		fmt.Println("")
	}
}
