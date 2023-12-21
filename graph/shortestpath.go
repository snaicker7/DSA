package main

import (
	"fmt"
	"google/graph/linkedlist"
)

type ItemDist struct {
	Node int
	Dist int
}
type Queue struct {
	Items []ItemDist
}

func (q *Queue) Enque(item ItemDist) {
	q.Items = append(q.Items, item)
}
func (q *Queue) Deque() ItemDist {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
func (q *Queue) Front() ItemDist {
	item := q.Items[0]
	return item
}
func (q *Queue) Rear() ItemDist {
	n := len(q.Items)
	item := q.Items[n-1]
	return item
}
func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2}, {0, 4}, {4, 3}, {4, 5}, {4, 6},
		{2, 5},
	}
	src := 1
	dest := 5

	graphLinked := linkedlist.LinkedListUnidirectional(n, edges)
	val := checkIfPathExist(graphLinked, src, dest)
	fmt.Println(val)
}

func checkIfPathExist(graphLinked linkedlist.Graph, src, dest int) int {
	if src == dest {
		return 0
	}
	que := &Queue{}
	visited := make(map[int]bool)
	if _, ok := graphLinked[src]; ok {
		srcItem := ItemDist{
			Node: src,
			Dist: 0,
		}
		que.Enque(srcItem)
		visited[src] = true
		temp := graphLinked[src].Head
		for temp != nil {
			item := ItemDist{
				Node: temp.Val,
				Dist: srcItem.Dist + 1,
			}
			que.Enque(item)
			temp = temp.Next
		}
		for len(que.Items) != 0 {
			next := que.Deque()
			if next.Node == dest {
				return next.Dist
			}
			if _, ok := graphLinked[next.Node]; ok {
				if _, ok := visited[next.Node]; !ok {
					visited[next.Node] = true
					temp := graphLinked[next.Node].Head
					for temp != nil {
						if _, ok := visited[temp.Val]; !ok {
							item := ItemDist{
								Node: temp.Val,
								Dist: next.Dist + 1,
							}
							que.Enque(item)

						}
						temp = temp.Next
					}
				}

			}
		}
	}
	return -1

}
