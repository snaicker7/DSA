package main

import "fmt"

// Find the min k elements
func main() {
	var arr = []int{7, 9, 2, 1, 10, 3, 5, 8, 12, 50}
	var final = make([]int, 0)
	var k = 5
	// max heapify k elements
	for i := 0; i < k; i++ {
		final = MaxheapifyDownUp(final, arr[i])
	}
	// for remaining elements ,if incoming elements is smaller than, delete the first element in the array
	for i := k; i < len(arr); i++ {
		if arr[i] < final[0] {
			final = final[1:]
			final = MaxheapifyDownUp(final, arr[i])
		}
	}
	fmt.Println("Final", final)
}

func MaxheapifyDownUp(arr []int, val int) []int {
	arr = append(arr, val)
	fmt.Println("Array before heapify", arr)
	if len(arr) <= 1 {
		fmt.Println("val: ", val)
		fmt.Println("Array ", arr)
		return arr
	}
	i := len(arr) - 1

	for i > 0 {
		parent := (i - 1) / 2
		if arr[i] > arr[parent] {
			swap(i, parent, arr)
			i = parent
		} else {
			break
		}
	}
	fmt.Println("val: ", val)
	fmt.Println("Array ", arr)
	fmt.Println("----------------------------------------")
	return arr
}

func swap(i, parent int, arr []int) {
	temp := arr[i]
	arr[i] = arr[parent]
	arr[parent] = temp
}
