package main

import (
	"fmt"
)

func bubbleSort(tosort []int) {
	size := len(tosort)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if tosort[j] < tosort[j-1] {
				tosort[j], tosort[j-1] = tosort[j-1], tosort[j]
			}
		}
	}
}

func main() {
	//unsorted := []int{1, 19, 3, 12, 5, 80, 99, 500}

	var n int
	fmt.Print("Enter size of array to sort : ")
	fmt.Scan(&n)
	fmt.Println("\nPlease add numbers to the array below")

	unsorted := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&unsorted[i])
	}
	//	fmt.Println(unsorted)

	fmt.Println("Un-sorted : ", unsorted)

	bubbleSort(unsorted)

	fmt.Println("Sorted : ", unsorted)
}
