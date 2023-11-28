package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	table := make([]int, size)

	for i := range table {

		table[i] = make([]int, size)

		for j := range table {

			table[i][j] = (j + 1) * (i + 1)
		}
	}
	return table
}
func main() {
	fmt.Print(MultiplicationTable(3))
}
