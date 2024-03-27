package main

import "io"

func ReadString(r io.Reader) (string, error) {
	var result string
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if n == 0 {
			break
		}
		result += string(buf[0])
	}

	return result, nil
}
