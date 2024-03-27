package main

import (
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	buffer := make([]byte, len(seq))
	index := 0

	for {
		n, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return false, err
		}

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if buffer[i] == seq[index] {
				index++
				if index == len(seq) {
					return true, nil
				}
			} else {
				index = 0
			}
		}
	}

	return false, nil
}
