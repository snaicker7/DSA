package main

import (
	"google/graph/linkedlist"
	"google/graph/queue"
)

func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3},
	}
	// graphMat := matrix.Matrixrep(n, edges)
	src := 0
	dest := 5

	graphLinked := linkedlist.LinkedListUnidirectional(n, edges)
	checkIfPathExist(graphLinked, src, dest)
}

func checkIfPathExist(graphLinked linkedlist.Graph, src, dest int) bool {
	if src == dest {
		return true
	}
	que := &queue.Queue{}
	visited := make(map[int]bool)
	if _, ok := graphLinked[src]; ok {
		que.Enque(src)
		visited[src] = true
		temp := graphLinked[src].Head
		for temp != nil {
			que.Enque(temp.Val)
			temp = temp.Next
		}
		for len(que.Items) != 0 {
			next := que.Deque()
			if next == dest {
				return true
			}
			if _, ok := graphLinked[next]; ok {
				if _, ok := visited[next]; !ok {
					visited[next] = true
					temp := graphLinked[next].Head
					for temp != nil {
						if _, ok := visited[temp.Val]; !ok {
							que.Enque(temp.Val)

						}
						temp = temp.Next
					}
				}

			}
		}
	}
	return false

}
