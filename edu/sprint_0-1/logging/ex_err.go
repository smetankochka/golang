package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Println("[ERROR] Something went wrong!")
}
