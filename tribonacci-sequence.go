package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	tribonacciSequence := signature[:]

	if n <= 3 {
		return tribonacciSequence[:n]
	}

	for i := 3; i < n; i++ {
		tribonacciSequence = append(tribonacciSequence, tribonacciSequence[i-3]+tribonacciSequence[i-2]+tribonacciSequence[i-1])
	}

	return tribonacciSequence
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
}
