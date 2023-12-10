package main

import "fmt"

func main() {
	arr := []int{20, 3, 11, 6, 7, 100, 20}

	// 20 3 4 6 7 100 9 35 88
	// 3 4 6 7 9 20 35 88 100
	// 0 1 2 3 4 5 	6 	7 8
	//

	/*
	   0  1  2  3  4  5    6
	   20 3, 4, 6, 7, 100, 20

	*/

	ans := quicksort(arr, 0, len(arr)-1)
	fmt.Println(ans)
}

func quicksort(arr []int, start, end int) []int {

	if start < end {
		part, arr := partition(arr, start, end)
		fmt.Println("arr after partition", arr)
		fmt.Println("--------------------")
		arr = quicksort(arr, start, part-1)
		arr = quicksort(arr, part+1, end)
	}
	return arr

}

func partition(arr []int, low, high int) (int, []int) {
	if len(arr) > 0 {
		i := low
		j := high
		pivot := low
		//pivot as the start eleme
		for i < j {
			for arr[i] <= arr[pivot] && i < high {
				i++
			}
			for arr[j] > arr[pivot] && j > low {
				j--
			}
			if i < j {
				arr = swap(arr, i, j)
				fmt.Println("array after swaps", arr)
			}
		}
		arr = swap(arr, pivot, j)
		fmt.Println("final pivot", arr)
		return i, arr
	}
	return -1, arr
}
func swap(arr []int, src, dest int) []int {
	temp := arr[src]
	arr[src] = arr[dest]
	arr[dest] = temp
	return arr
}
