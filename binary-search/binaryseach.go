package main

import "fmt"

//Logic:

// check if start <= end,
// 			if yes check if element is present at the mid element
//						 if no, then check if the element> mid element, change the start
//										   if no, check change the end
func main() {
	arr := []int{1, 3, 5, 7, 8, 11, 100}
	val := 3
	ans := binarysearch(arr, 0, len(arr)-1, val)
	fmt.Println(ans)
}

func binarysearch(arr []int, start, end, val int) int {
	if start <= end {
		//mid := (start + end) / 2
		mid := start + ((end - start) / 2) // To prevent int overflow if start & end both are 2^31 -1
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
