package main

import (
	"os"
)

func ModifyFile(filename string, pos int64, val string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Seek(pos, 0); err != nil {
		return err
	}
	if _, err := file.Write([]byte(val)); err != nil {
		return err
	}
	return nil
}
