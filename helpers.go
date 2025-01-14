package main

// Factorial calculates the factorial of a given non-negative integer n.
// If n is 0, it returns 1 as 0! is 1.
// Otherwise, it returns n * Factorial(n-1).
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
