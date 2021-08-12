package main

import "fmt"

//二分查找
func BinarySearch(arr []int, left int, right int, target int) int {
	if left > right {
		return -1 //查找失败
	}
	middle := left + (right-left)/2
	if arr[middle] == target {
		return middle
	} else if arr[middle] < target {
		return BinarySearch(arr, middle+1, right, target)
	} else {
		return BinarySearch(arr, left, middle-1, target)
	}
}
func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := BinarySearch(number, 0, 9, 6)
	fmt.Println(index)

}
