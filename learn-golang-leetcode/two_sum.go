package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			// fmt.Println(i, j)
			fmt.Println(nums[i], nums[j])
			if i != j {
				// fmt.Println(nums[i], nums[j])
				sum := nums[i] + nums[j]
				if sum == target {
					result = []int{j, i}
				}
			}
		}
	}

	return result
}

func main() {
	// intData := []int{2, 7, 8, 9}
	intData2 := []int{3, 2, 3}

	result := twoSum(intData2, 6)

	fmt.Println(result)
}
