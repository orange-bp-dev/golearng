package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("Hello"[0]) // 72
	fmt.Println(string("Hello"[0])) // H

	var s string = "Hello"
	fmt.Println(strings.Replace(s, "H", "X", 1)) // Xello
	fmt.Println(strings.Contains(s, "He")) // true
	fmt.Println(`Test
										TEST
      てすと	`)
fmt.Println("\"")
fmt.Println(`"`)
}