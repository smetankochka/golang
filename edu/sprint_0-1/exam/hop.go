package main

import "fmt"
import "math"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 3; i <= maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int
	answer := ""
	fmt.Scan(&a)
	for i := 3; i <= a; i += 5 {
		if isPrime(i) {
			answer += "хоп "
		} else {
			answer += fmt.Sprintf("%d ", i)
		}
	}
	fmt.Println(answer)
}
