package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 333}

	i, ok := binarySearch(arr, 4)
	if ok {
		fmt.Println(i)
	}
}

func binarySearch(arr []int, target int) (int, bool) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == target {
			return middle, true
		} else if target > arr[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return 0, false
}
