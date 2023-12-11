package main

import "fmt"

func insertion(arr []int, val int) []int {
	arr = append(arr, val)
	child := len(arr) - 1

	for child > 0 {
		par := (child - 1) / 2
		if arr[child] < arr[par] {
			arr[child], arr[par] = arr[par], arr[child]
			child = par
		}
	}
	return arr
}
func deletion(arr []int) ([]int, int) {
	fmt.Println("before deletion", arr)
	root := peek(arr)
	n := len(arr) - 1
	lastElem := arr[n]
	arr[0] = lastElem
	arr = arr[:n]
	fmt.Println("After deletion", arr)
	for i := 0; i < n/2; {
		min := i
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if arr[leftchild] < arr[i] {
			min = leftchild
		}
		if arr[rightchild] < arr[min] {
			min = rightchild
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
			i = min
		} else {
			return arr, root
		}
	}
	return arr, root
}

func peek(arr []int) int {
	return arr[0]
}

func main() {
	arr := []int{3, 5, 9, 7, 14, 11, 13}
	arr, root := deletion(arr)
	fmt.Print("root", root)
	fmt.Print("arr", arr)
	arr = insertion(arr, 1)
	fmt.Print("insertion arr", arr)
}
