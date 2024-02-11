package main

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N == 2 || N == 1 {
		fmt.Println("YES")
	} else if N%2 == 1 {
		fmt.Println("NO")
	} else {
		IsPowerOfTwoRecursive(N / 2)
	}
}
