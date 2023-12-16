package main

import "fmt"

func main() {
	subsetArr := []int{1, 2, 3}
	var arr = make([]int, 0)
	var result [][]int
	generate(0, arr, subsetArr, &result)

}

//subarray

//	func generate(curInd int, arr, subsetArr []int) {
//		if curInd == len(subsetArr) {
//			return
//		}
//		arr = append(arr, subsetArr[curInd])
//		fmt.Println(arr)
//		generate(curInd+1, arr, subsetArr)
//	}
func generate(curInd int, arr, subsetArr []int, result *[][]int) {
	if curInd >= len(subsetArr) {
		fmt.Println(arr)
		*result = append(*result, arr)
		return
	}

	arr = append(arr, subsetArr[curInd])
	generate(curInd+1, arr, subsetArr, result)
	arr = arr[:len(arr)-1]
	generate(curInd+1, arr, subsetArr, result)
}
