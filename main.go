/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.*/
package main

import "fmt"

func main() {

	slice := []int{2, 4, 5, 7, 8, 13, 45, 67}
	r := twoSum(slice, 6)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {

		if v, found := m[target-num]; found {
			return []int{v, idx}
		}

		m[num] = idx
	}
	return nil
}
