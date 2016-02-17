package main

import "fmt"

func main() {
	arrays := [][]int{
		[]int{2, 2, 3, 4, 4},
		// assume the duplicate can happen multiple times
		[]int{3, 4, 4, 7, 2, 1, 2, 3, 3, 1, 1, 7, 5, 8, 8},
	}

	for _, a := range arrays {
		unique := findUnique(a)
		fmt.Println("Given ", a)
		fmt.Println("  => Found unique int: ", unique)
	}
}

func findUnique(a []int) int {
	dups := make(map[int]int, 0)

	for _, i := range a {
		if dups[i] == 0 {
			dups[i] = 1
		} else {
			dups[i]++
		}
	}

	for k, v := range dups {
		if v == 1 {
			return k
		}
	}

	panic("no unique int found!")
}
