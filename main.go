/*Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.*/
package main

import "fmt"

func main() {

	slice := []int{2, 4, 5, 7, 8, 13, 45, 67}
	r := searchInsert(slice, 3)
	fmt.Println(r)
}
func searchInsert(nums []int, target int) int {
	i := 0

	for i < len(nums) && nums[i] < target {
		i++
	}

	return i
}
