package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 1, 2, 4}

	res := sumSubarrayMins(arr)
	fmt.Println(res)
}

func sumSubarrayMins(arr []int) int {
	l := len(arr)
	sum := 0

	for i := 0; i < l; i++ {
		prevMin := arr[i]
		sum += prevMin

		for j := i + 1; j != l; j++ {
			if prevMin > arr[j] {
				prevMin = arr[j]
			}
			sum += prevMin
		}
	}

	return sum
}

//Input: arr = [3,1,2,4]
//Output: 17
//Explanation:
//Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].

// sum [3], [1], [2], [4]
// remainder [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]

// sum Min[3,1] Min[3,1,2] Min[3,1,2,4]
// remainder [1,2], [2,4], [1,2,4]

// sum Min[1,2], Min[1,2,4]
// remainder [2,4]

// sum Min[2,4]
// remainder []
