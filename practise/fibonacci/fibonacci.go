package fibonacci

func FibonacciWithoutMemoization(n int) int {
	if n <= 1 {
		return 1
	}
	return FibonacciWithoutMemoization(n-1) + FibonacciWithoutMemoization(n-2)
}

func FibonacciWithMemoization(n int, fibMap map[int]int) int {
	if n <= 1 {
		return 1
	}
	if fibMap[n] != 0 {
		return fibMap[n]
	} else {
		fibMap[n] = FibonacciWithMemoization(n-1, fibMap) + FibonacciWithMemoization(n-2, fibMap)
	}

	return fibMap[n]
}
