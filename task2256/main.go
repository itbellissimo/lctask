package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 5, 3, 9, 5, 3}

	fmt.Println(nums)
	res := minimumAverageDifference(nums)
	fmt.Println(res)
}

func minimumAverageDifference(nums []int) int {
	var (
		frontSum   = 0
		endSum     = 0
		minAvgDiff = math.MaxInt
		index      = 0
	)

	for i := 0; i < len(nums); i++ {
		endSum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		frontSum += nums[i]
		endSum -= nums[i]

		frontAvg := float64(frontSum / (i + 1))

		endAvg := 0.0
		if i < len(nums)-1 {
			endAvg = float64(endSum / (len(nums) - i - 1))
		}

		if m := math.Abs(frontAvg - endAvg); int(math.Floor(m)) < minAvgDiff {
			minAvgDiff = int(math.Floor(m))
			index = i
		}
	}

	return index
}
