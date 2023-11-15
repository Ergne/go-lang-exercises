package main

import (
	"fmt"
)

func Solution(number int) string {
	result := ""

	// Use a slice instead of an array for roman
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	decimal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for index, value := range decimal {
		for number >= value {
			number -= value
			result += roman[index]
		}
	}
	return result
}

func main() {
	fmt.Println(Solution(99))
}
