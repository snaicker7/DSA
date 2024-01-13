package main

import "fmt"

/*
0->1, 2
1->0
2->0
3->4, 5
4->3, 5
5->3, 4
*/
func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3},
	}
	graph := LinkedListUnidirectional(n, edges)
	graph.dfsfunc()
}
func (nodes Graph) dfsfunc() {
	visited := make(map[int]bool)
	var connectedcomp int = 0
	for node, list := range nodes {
		if _, ok := visited[node]; !ok {
			connectedcomp++
			visited[node] = true
			fmt.Println(node)
			nodes.dfs(list, visited)
		}
	}
	fmt.Println(connectedcomp)
}

func (nodes Graph) dfs(node List, visited map[int]bool) {
	temp := node.Head
	for temp != nil {
		if _, ok := visited[temp.Val]; !ok {
			visited[temp.Val] = true
			fmt.Println(temp.Val)
			nodes.dfs(nodes[temp.Val], visited)
		}
		temp = temp.Next
	}

}

type Node struct {
	Val  int
	Next *Node
}
type List struct {
	Head *Node
}
type Graph map[int]List

func (linkedlistmap Graph) LinkedListInsertion(i int, node *Node) {
	if _, ok := linkedlistmap[i]; !ok {
		linkedlistmap[i] = List{
			Head: node,
		}
	} else {
		temp := linkedlistmap[i].Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
	}
}
func LinkedListUnidirectional(n int, edges [][]int) Graph {
	var linkedlistmap = make(Graph)
	for _, edge := range edges {
		i := edge[0]
		j := edge[1]
		node := &Node{
			Val:  j,
			Next: nil,
		}
		linkedlistmap.LinkedListInsertion(i, node)
		node = &Node{
			Val:  i,
			Next: nil,
		}
		linkedlistmap.LinkedListInsertion(j, node)
	}
	return linkedlistmap
}
