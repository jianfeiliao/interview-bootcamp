package main

import "fmt"

func main() {
	arrays := [][]int{
		[]int{},
		[]int{2},
		[]int{5, 2},
		[]int{5, 2, 3, 1},
		[]int{3, 1, 4, 2, 0, 9, 8, 5, 6},
		[]int{3, 3, 4, 2, 2, 5, 8, 5, 7},
	}

	for _, a := range arrays {
		fmt.Println("unsort: ", a)
		QuickSort(a)
		fmt.Println("sorted: ", a)
		fmt.Println()
	}

}

func QuickSort(a []int) {
	quickSortHelper(a, 0, len(a)-1)
}

func partition(a []int, start, end, pivotIndex int) int {
	//fmt.Println("partition start:", a, start, end, a[pivotIndex])
	pivotValue := a[pivotIndex]
	swap(a, pivotIndex, end) // move pivot to the end
	correctPivotIndex := start
	for i := start; i < end; i++ {
		//fmt.Println("i=", i, a[i], "<", pivotValue, "?")
		if a[i] < pivotValue {
			swap(a, correctPivotIndex, i)
			correctPivotIndex++
		}
	}
	swap(a, end, correctPivotIndex) // move pivot to the correct place
	//fmt.Println("partition done: ", a, start, end, correctPivotIndex)
	return correctPivotIndex
}

func quickSortHelper(a []int, start, end int) {
	if start >= end {
		return
	}

	correctPivotIndex := partition(a, start, end, (start+end)/2)

	quickSortHelper(a, start, correctPivotIndex-1)
	quickSortHelper(a, correctPivotIndex+1, end)
}

func swap(a []int, x, y int) {
	//fmt.Println("swap", a[x], "and", a[y])
	temp := a[x]
	a[x] = a[y]
	a[y] = temp
	//fmt.Println("   =>", a)
}
