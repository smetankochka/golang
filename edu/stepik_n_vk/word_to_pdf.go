package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Путь к исходному Word-документу
	inputFilePath := "/home/smetanka/Загрузки/Telegram Desktop/name_of_react.docx"
	// Путь к целевому PDF-файлу
	outputFilePath := "/home/smetanka/Загрузки/Telegram Desktop/name_of_react_2.pdf"
	// Команда для вызова unoconv
	cmd := exec.Command("unoconv", "-f", "pdf", "-o", outputFilePath, inputFilePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Документ успешно сконвертирован в PDF.")
}
