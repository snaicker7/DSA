package main

import "fmt"

func main() {
	var maxleftheap []int
	var minrightheap []int
	var medianArray []int
	var median = 0
	var input = []int{5, 17, 100, 11}
	// The idea is splitting the array into 1 equal parts such that
	// 1st half will have max heap of smaller elements
	// 2nd half will have a min heap of larger elements
	//5, 11----17, 100-> in this example, median will be 11+17/2
	// for upcoming streams, we will check if the value is < median, add it in the max heap at the first heap
	// if val > median, then add it min heap at the 2nd half of the array
	for _, val := range input {
		fmt.Println("max left heap", maxleftheap)
		fmt.Println("min right heap", minrightheap)
		fmt.Println("median", median)
		fmt.Println("val", val)
		maxleftheap, minrightheap, median = findMedian(maxleftheap, minrightheap, val, median)
		medianArray = append(medianArray, median)
		fmt.Println("max left heap", maxleftheap)
		fmt.Println("min right heap", minrightheap)
		fmt.Println("median array", medianArray)
		fmt.Println("--------------------------------------------------------------")

	}

}
func findMedian(maxleftheap []int, minrightheap []int, val int, median int) ([]int, []int, int) {
	if len(maxleftheap) == 0 {
		maxleftheap = append(maxleftheap, val)
		return maxleftheap, minrightheap, peek(maxleftheap)
	}

	// if the value is less than the median, then add it to maxleftheap
	if val < median {
		if len(maxleftheap)-len(minrightheap) >= 1 {
			peekmax := peek(maxleftheap)

			minrightheap = heapifyMin(minrightheap, peekmax)
			maxleftheap = delete(maxleftheap)
			maxleftheap = heapifyMax(maxleftheap, val)
		} else {
			maxleftheap = heapifyMax(maxleftheap, val)
		}
	} else {
		if len(minrightheap)-len(maxleftheap) >= 1 {
			peekmin := peek(minrightheap)
			maxleftheap = heapifyMax(maxleftheap, peekmin)
			minrightheap = delete(minrightheap)
			minrightheap = heapifyMin(minrightheap, val)

		} else {
			minrightheap = heapifyMin(minrightheap, val)

		}

	}
	if len(maxleftheap) == len(minrightheap) {
		median = (peek(maxleftheap) + peek(minrightheap)) / 2
	} else if len(maxleftheap) > len(minrightheap) {
		median = peek(maxleftheap)
	} else {
		median = peek(minrightheap)
	}
	return maxleftheap, minrightheap, median
}

func peek(heap []int) int {
	if len(heap) > 0 {
		return heap[0]
	}
	return 0

}
func delete(heap []int) []int {
	return heap[1:]
}
func heapifyMin(heap []int, val int) []int {
	heap = append(heap, val)
	i := len(heap) - 1

	for i > 0 {
		par := (i - 1) / 2
		//only this line changes
		if heap[i] < heap[par] {
			swap(heap, i, par)
			i = par
		} else {
			return heap
		}
	}
	return heap
}
func swap(heap []int, i, par int) []int {
	temp := heap[i]
	heap[i] = heap[par]
	heap[par] = temp
	return heap
}
func heapifyMax(heap []int, val int) []int {
	heap = append(heap, val)
	i := len(heap) - 1

	for i > 0 {
		par := (i - 1) / 2
		//only this line changes
		if heap[i] > heap[par] {
			swap(heap, i, par)
			i = par
		} else {
			return heap
		}
	}
	return heap
}
