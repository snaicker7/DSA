package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 5, 6}
	arr2 := []int{3, 5, 7, 9}
	k := 20
	fmt.Println(ksmallestsumpair(arr1, arr2, k))
}

type Arrind struct {
	sum     int
	arr1ind int
	arr2ind int
}

//insertion into heap-> bottom up approach
//deletion from a heap-> top down approach
func ksmallestsumpair(arr1, arr2 []int, k int) [][]int {
	var final = make([][]int, 0)
	var indexmap = make(map[string]struct{})
	i := 0
	j := 0
	indexStr := fmt.Sprintf("%d-%d", i, j)
	indexmap[indexStr] = struct{}{}
	var sumarr = []Arrind{}
	var smallestsum = Arrind{}

	smallestsum = Arrind{
		sum:     arr1[i] + arr2[j],
		arr1ind: i,
		arr2ind: j,
	}
	i = 0
	j = 0
	final = append(final, []int{arr1[smallestsum.arr1ind], arr2[smallestsum.arr2ind]})
	for s := 2; s <= k && s <= len(arr1)*len(arr2); {
		var val1 = Arrind{}
		var val2 = Arrind{}
		if i+1 < len(arr1) {
			val1 = Arrind{
				sum:     arr1[i+1] + arr2[j],
				arr1ind: i + 1,
				arr2ind: j,
			}
			indexStr := fmt.Sprintf("%d-%d", val1.arr1ind, val1.arr2ind)
			if _, ok := indexmap[indexStr]; !ok {
				sumarr = summininsertion(sumarr, val1)
				indexmap[indexStr] = struct{}{}
			}

		}
		if j+1 < len(arr2) {
			val2 = Arrind{
				sum:     arr1[i] + arr2[j+1],
				arr1ind: i,
				arr2ind: j + 1,
			}
			indexStr := fmt.Sprintf("%d-%d", val2.arr1ind, val2.arr2ind)
			if _, ok := indexmap[indexStr]; !ok {
				sumarr = summininsertion(sumarr, val2)
				indexmap[indexStr] = struct{}{}
			}

		}
		if len(sumarr) > 0 {
			sumarr, smallestsum = summindeletion(sumarr, len(sumarr)-1)
			indexStr1 := fmt.Sprintf("%d-%d", smallestsum.arr1ind+1, smallestsum.arr2ind)
			indexStr2 := fmt.Sprintf("%d-%d", smallestsum.arr1ind, smallestsum.arr2ind+1)
			if _, ok := indexmap[indexStr1]; !ok {
				i = smallestsum.arr1ind
				j = smallestsum.arr2ind
				s++
			} else if _, ok := indexmap[indexStr2]; !ok {
				i = smallestsum.arr1ind
				j = smallestsum.arr2ind
				s++
			}

			final = append(final, []int{arr1[smallestsum.arr1ind], arr2[smallestsum.arr2ind]})
		} else {
			break
		}

	}
	return final

}
func summindeletion(arr []Arrind, n int) ([]Arrind, Arrind) {
	if len(arr) > 0 {
		root := peek(arr)
		lastElem := arr[n]
		arr[0] = lastElem
		arr = arr[:n]
		n = len(arr) - 1
		for i := 0; i <= n/2; {
			min := i
			leftchild := 2*i + 1
			rightchild := 2*i + 2
			if leftchild <= n && arr[leftchild].sum < arr[i].sum {
				min = leftchild
			}
			if rightchild <= n && arr[rightchild].sum < arr[min].sum {
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

	return arr, Arrind{}
}
func peek(arr []Arrind) Arrind {
	return arr[0]
}
func summininsertion(arr []Arrind, val Arrind) []Arrind {
	arr = append(arr, val)
	child := len(arr) - 1

	for child > 0 {
		par := (child - 1) / 2
		if arr[child].sum < arr[par].sum {
			arr[child], arr[par] = arr[par], arr[child]
			child = par
		} else {
			return arr
		}
	}
	return arr
}
