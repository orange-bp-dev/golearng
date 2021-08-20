package main

import (
	"fmt"
	"os/user"
	"time"
)

//ターミナル上でGoコマンドの使い方を調べる
//godoc fmt Println

func main() {
	fmt.Println("Hi", time.Now())
	fmt.Println(user.Current())
}