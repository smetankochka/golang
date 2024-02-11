package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	letter1, _ := r.ReadString('\n')
	letter1 = strings.TrimSpace(letter1)
	letter2, _ := r.ReadString('\n')
	letter2 = strings.TrimSpace(letter2)
	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)
	c1 := strings.Count(text, letter1)
	c2 := strings.Count(text, letter2)
	if c1 < c2 {
		fmt.Printf("%s %d\n", letter2, c2)
		fmt.Printf("%s %d", letter1, c1)
	} else {
		fmt.Printf("%s %d\n", letter1, c1)
		fmt.Printf("%s %d", letter2, c2)
	}
}
