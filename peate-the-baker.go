package main

import (
	"fmt"
	"math"
)

func Cakes(recipe, available map[string]int) int {
	maxCakes := math.MaxInt
	for key, value := range recipe {
		if available[key]/value < maxCakes {
			maxCakes = available[key] / value
		}
	}
	return maxCakes
}

func main() {
	fmt.Println(Cakes(
		map[string]int{"oil": 82, "pears": 92},
		map[string]int{"chocolate": 614, "flour": 1270, "nuts": 2237, "pears": 258}))
}
