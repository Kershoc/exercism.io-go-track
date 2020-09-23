package diffsquares

//SquareOfSum returns (1 + 2 + ... + n)²
func SquareOfSum(n int) (sum int) {
	// [n(n+1)/2]^2
	sum = n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares returns 1² + 2² + ... + n²
func SumOfSquares(n int) int {
	// Σn2 = [n(n+1)(2n+1)]/6
	return (n * (n + 1) * (2*n + 1)) / 6
}

//Difference returns (1 + 2 + ... + n)² - (1² + 2² + ... + n²)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
