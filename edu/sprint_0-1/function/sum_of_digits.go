package main

func SumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + SumDigitsRecursive((n-n%10)/10)
}
