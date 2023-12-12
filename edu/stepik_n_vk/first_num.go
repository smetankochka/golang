package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    for a > 9 {
        a /= 10
    }
    fmt.Print(a)
}
