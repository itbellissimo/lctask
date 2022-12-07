package main

import (
	"fmt"
	"math"
)

func main() {
	area := 122122

	fmt.Println(area)
	res := constructRectangle(area)
	fmt.Println(res)
}

func constructRectangle(area int) []int {
	for y := int(math.Sqrt(float64(area))); ; y-- {
		if area%y == 0 {
			return []int{area / y, y}
		}
	}
}
