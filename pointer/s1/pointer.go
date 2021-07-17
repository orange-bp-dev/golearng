package main

import "fmt"

func main() {
	var n int = 100

	//nの値
	fmt.Println(n)

	//nを入れたメモリのアドレス（16進数）
	fmt.Println(&n)

	//nのポインタ型
	var p *int = &n

	//アドレスが表示される
	fmt.Println(p)

	//アドレスが指しているメモリの中身が表示される
	fmt.Println(*p)
}