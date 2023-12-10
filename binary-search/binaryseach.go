package main

import "fmt"

//Logic:

// check if start <= end, if yes, then check if the element
func main() {
	arr := []int{1, 3, 5, 7, 8, 11, 100}
	val := 3
	ans := binarysearch(arr, 0, len(arr)-1, val)
	fmt.Println(ans)
}

func binarysearch(arr []int, start, end, val int) int {
	if start <= end {
		mid := (start + end) / 2
		if arr[mid] == val {
			return val
		} else {
			if val < arr[mid] {
				return binarysearch(arr, start, mid-1, val)
			} else {
				return binarysearch(arr, mid+1, end, val)
			}
		}
	}
	return -1

}
