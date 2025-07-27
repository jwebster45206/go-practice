package main

import "time"

func main() {
	// Example usage
	startMS := time.Now()
	println(factorialRecursive(100)) // Output: 120
	duration := time.Since(startMS).Microseconds()
	println("Execution time (recursive):", duration, "microseconds")

	startMS = time.Now()
	println(factorialIterative(100)) // Output: 120
	duration = time.Since(startMS).Microseconds()
	println("Execution time (iterative):", duration, "microseconds")
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
