package main

func CalculateDigitalRoot(n int) int {
	for n > 9 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	return n
}
