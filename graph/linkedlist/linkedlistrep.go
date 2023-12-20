package linkedlist

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
