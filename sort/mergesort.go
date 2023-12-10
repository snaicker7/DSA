package main

import "fmt"

func main() {
	arr := []int{20, 3, 4, 6, 7, 100, 20}
	ans := mergesort(arr)
	fmt.Println(ans)
}

func mergesort(arr []int) []int {

	// find the mid, and divide the array into left sub array and right sub array
	// merge the left and right sub array
	// left = 0-n/2
	// right= n/2+1-last index
	start := 0
	end := len(arr) - 1
	if start < end {
		mid := start + ((end - start) / 2)
		leftsubarray := mergesort(arr[start : mid+1])
		rightsubarray := mergesort(arr[mid+1 : end+1])
		return merge(leftsubarray, rightsubarray)
	}
	return arr

}

func merge(leftsubarr, rightsubarray []int) []int {
	var finalarr = make([]int, 0)
	i := 0
	j := 0
	for i < len(leftsubarr) && j < len(rightsubarray) {
		if leftsubarr[i] < rightsubarray[j] {
			finalarr = append(finalarr, leftsubarr[i])
			i++
		} else {
			finalarr = append(finalarr, rightsubarray[j])
			j++
		}
	}
	if i < len(leftsubarr) {

		finalarr = append(finalarr, leftsubarr[i:]...)
	}
	if j < len(rightsubarray) {

		finalarr = append(finalarr, rightsubarray[j:]...)
	}
	return finalarr
}
