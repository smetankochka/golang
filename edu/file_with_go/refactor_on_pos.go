package main

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Seek(int64(pos), 0); err != nil {
		return err
	}
	if _, err := file.Write([]byte(val)); err != nil {
		return err
	}
	return nil
}

// В этой функции мы открываем файл с помощью os.OpenFile с
// флагом os.O_RDWR, что позволяет нам читать и записывать данные.
// Затем мы перемещаемся к указанной позиции в файле с использованием file.Seek,
// и, наконец, записываем значение val в эту позицию с помощью file.Write.
