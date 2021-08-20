package main

import (
	"fmt"
	// "mylib"
)

/*
main()よりも先に実行される特別な関数
初期設定などで使われる
*/
func init() {
	// mylib.Average(1)
	fmt.Println("Init!")
}

//main()内で実行する必要あり
func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	 //大文字でPrint
	fmt.Println("Hello world!", "TEST")
}