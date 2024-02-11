package main

type Number interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T Number](arr []T) T {
	var total T
	for _, n := range arr {
		total += n
	}
	return total
}
