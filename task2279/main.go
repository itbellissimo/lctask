package main

import "sort"

func main() {

}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	if len(capacity) != len(rocks) {
		return 0
	}
	diff := make([]int, len(capacity))
	for i := 0; i < len(capacity); i++ {
		diff[i] = capacity[i] - rocks[i]
	}

	sort.Ints(diff)

	countBugs := 0
	i := 0
	for additionalRocks > 0 {
		additionalRocks = additionalRocks - diff[i]
		if additionalRocks >= 0 {
			countBugs++
		}
		i++
		if i >= len(capacity) {
			break
		}
	}

	return countBugs
}
