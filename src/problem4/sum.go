package main

func main() {
	assert(sum_to_n_a(10) == 55, "sum_to_n_a(10) should return 55")
	assert(sum_to_n_b(100) == 5050, "sum_to_n_b(100) should return 5050")
	assert(sum_to_n_c(1000) == 500500, "sum_to_n_c(1000) should return 500500")
}

// sum_to_n_a is the for-loop implementation of sum_to_n
// will take O(n) time
func sum_to_n_a(n int) int {
	var sum = 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// sum_to_n_b is the recursive implementation of sum_to_n
// will take O(n) time
func sum_to_n_b(n int) int {
	if n == 0 {
		return 0
	}

	return sum_to_n_b(n-1) + n
}

// sum_to_n_c is the formula-based implementation of sum_to_n
// use the identity 1 + 2 + ... + n = n(n + 1)/2
// will take O(1) time, most efficient
func sum_to_n_c(n int) int {
	return n * (n + 1) / 2
}

// asssert function will check if a condition is satisfied
func assert(condition bool, message string) {
	if !condition {
		panic("Assertion failed: " + message)
	}
}
