package main

import (
	"strings"
)

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(data []byte) (int, error) {
	str := strings.ToUpper(string(data))
	u.UpperString = str
	return len(data), nil
}
