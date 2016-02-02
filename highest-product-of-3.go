package main

import (
	"fmt"
	"math"
)

func main() {
	arrays := [][]int{
		[]int{4, 3, 2, 6},
		[]int{-3, -1, -6},
		[]int{-10, -10, 1, 3, 2},
		[]int{4, 3, 2, 1, 2, -3},
		[]int{4, -6, 2, 5, -2, -4},
	}

	for _, a := range arrays {
		fmt.Println(a, "=>", highestProduct(a))
	}
}

// bruteforce
func highestProduct(a []int) int {
	if len(a) < 3 {
		panic("need at least 3 number")
	}

	max := math.MinInt64

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				if i != j && i != k && j != k {
					p3 := a[i] * a[j] * a[k]
					if p3 >= max {
						max = p3
					}
				}
			}
		}
	}
	return max
}
