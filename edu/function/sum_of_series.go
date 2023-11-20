package main

func CalculateSeriesSum(n int) float64 {
	if n == 1 {
		return 1
	}
	return 1 / float64(n) + CalculateSeriesSum(n-1)
}
