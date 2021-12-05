package main

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6

func sumOfFirst(n int) int {
	var num int
	for num = n; n > 0; num -= 1 {
		num = num + 1
	}
	return num
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	var num int = n % 1
	var i int
	var result bool
	for i = 0; i < n; i++ {
		if num == 0 && num == n {
			result = true
		} else {
			result = false
		}
	}
	return result
}
