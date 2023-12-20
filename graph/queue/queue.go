package queue

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) Enque(item int) {
	q.Items = append(q.Items, item)
}
func (q *Queue) Deque() int {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
func (q *Queue) Front() int {
	item := q.Items[0]
	return item
}
func (q *Queue) Rear() int {
	n := len(q.Items)
	item := q.Items[n-1]
	return item
}
func main() {
	qu := &Queue{}
	qu.Enque(1)
	qu.Enque(2)
	qu.Enque(3)
	qu.Deque()
	fmt.Print("qu", qu.Items)
}
