package main

import (
	"fmt"
	"os/user"
)

func init() {
	fmt.Println("init")
}

func main() {
	bazz()
	fmt.Println("Hello world")
}

func bazz() {
	fmt.Println(user.Lookup("satoryoka"))
}
