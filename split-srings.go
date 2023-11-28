package main

import (
	"fmt"
	"math"
)

func Solution(str string) []string {
	result := make([]string, int(math.Ceil(float64(len(str))/2)))
	for i, char := range str {
		result[i/2] += string(char)
	}
	if len(result[len(result)-1]) < 2 {
		result[len(result)-1] += "_"
	}
	return result
}

func main() {
	fmt.Println(Solution("abcdddddde"))
}
