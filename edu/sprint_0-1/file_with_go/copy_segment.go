package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	file, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Seek(int64(startpos), 0)
	file2, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer file2.Close()
	buffer := make([]byte, 100)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			return nil
		} // the end of the file
		if err != nil {
			return err
		}
		_, err = file2.Write(buffer[:n])
		if err != nil {
			return err
		}
	}
	return nil
}
