package main

import "fmt"

func main() {
	//空変数用のメモリを確保
	var p *int = new(int)

	//メモリのアドレスが表示
	fmt.Println(p)

	//アドレスの中身が表示
	fmt.Println(*p) // 0
	*p++
	fmt.Println(*p) // 1

	//変数宣言してるがメモリを確保していない！
	var p2 *int
	fmt.Println(p2) //<nil>


	// makeとnewの違いは？？
	// ポインタを返す変数はnewそれ以外はmakeをつかう！
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)
	
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

}