package main

import "fmt"

func main() {
	subsetArr := []int{1, 2, 3}
	var arr = make([]int, 0)
	generate(0, arr, subsetArr)

}
func generate(curInd int, arr, subsetArr []int) {
	if curInd == len(subsetArr) {
		return
	}
	arr = append(arr, subsetArr[curInd])
	fmt.Println(arr)
	generate(curInd+1, arr, subsetArr)
}
