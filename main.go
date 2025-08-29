package main

import "fmt"

func binarySearch(arr []int, t int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == t {
			return mid
		} else if arr[mid] < t {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {

	a := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	fmt.Println(a)

}
