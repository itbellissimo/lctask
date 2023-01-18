package main

import "math"

func main() {

}

// maxIceCream
// https://leetcode.com/problems/maximum-ice-cream-bars/solutions/2885725/maximum-ice-cream-bars/?orderBy=most_votes
func maxIceCream(costs []int, coins int) int {
	costsCount := make(map[int]int, 1)
	maxCost := 0
	iceCreamCount := 0
	for i := range costs {
		if maxCost < costs[i] {
			maxCost = costs[i]
		}
		costsCount[costs[i]]++
	}

	for cost := 1; cost <= maxCost; cost++ {
		if _, ok := costsCount[cost]; !ok {
			continue
		}

		if coins < cost {
			break
		}

		c := int(math.Floor(math.Min(float64(costsCount[cost]), math.Floor(float64(coins/cost)))))

		coins -= cost * c

		iceCreamCount += c
	}

	return iceCreamCount
}
