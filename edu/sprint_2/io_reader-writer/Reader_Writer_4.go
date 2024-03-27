package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, n)
	totalRead := uint(0)

	for totalRead < n {
		toRead := n - totalRead
		if toRead > uint(len(data)) {
			toRead = uint(len(data))
		}

		read, err := r.Read(data[:toRead])
		if err != nil && err != io.EOF {
			return err
		}

		if read == 0 {
			break
		}

		written, err := w.Write(data[:read])
		if err != nil {
			return err
		}

		totalRead += uint(written)
		if read < len(data) {
			break
		}
	}

	return nil
}
