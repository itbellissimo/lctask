package main

import (
	"fmt"
)

func main() {
	var nums = []int{0, 0, 0}

	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	j := len(nums) - 1
	for i := 0; i <= j; i++ {
		if nums[i] != 0 {
			break
		}
		if i == j {
			return
		}
	}
	for i := 0; i < j; i++ {
		if nums[i] == 0 {
			k := i + 1
			for nums[k] == 0 && k < j {
				k++
			}
			if k > j {
				break
			}
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
}
