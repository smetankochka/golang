package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var ansarr []T
	for _, x := range arr {
		if predicate(x) {
			ansarr = append(ansarr, x)
		}
	}
	return ansarr
}
