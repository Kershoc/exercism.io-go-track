package diffsquares

//SquareOfSum returns (1 + 2 + ... + n)²
func SquareOfSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	sum = sum * sum
	return
}

//SumOfSquares returns 1² + 2² + ... + n²
func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return
}

//Difference returns (1 + 2 + ... + n)² - (1² + 2² + ... + n²)
func Difference(n int) int {
	soq := 0
	qos := 0
	for i := 1; i <= n; i++ {
		soq += i
		qos += i * i
	}

	return soq*soq - qos
}
