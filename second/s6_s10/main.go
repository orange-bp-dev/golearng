package main

import (
	"fmt"
	"os/user"
)

// 別ファイルから呼び出される場合は変数は大文字スタート!!
const Pi = 3.14

// const では型を定義しない！
// const では変数実行時に型推測！！
const (
	Username = "Sato"
	Pass = "Ryoka"
)

func init() {
	fmt.Println("init")
}

func main() {
	// 同じ型の場合はカンマで区切れば同時に宣言可能！
	var t, f bool = true, false
	var (
		i int
		f64 float64
	)
	// 通常は := で宣言するので不要
	// この書き方は関数内でしかかけない!!
	xt, xf := true, false


	bazz()
	//booleanの初期値はfalse
	//宣言だけして値を宣言しない場合、各型の初期値が入る
	fmt.Println(t, f, i, f64, xt, xf)

	// xf64 := 1.2

	//明示的に型をしてするにはvarが必要
	//float32 なにも宣言しない場合float64になってしまう
	var xf32 float32 = 1.2
	
	//%Tで型(Type)確認
	//%vで値(value)確認
	fmt.Printf("%T\n %v", xf32, xf32)

	//\nで改行
	fmt.Printf("%T\n", t)
	
	fmt.Println(10/3) //3(小数点以下切り捨て)
	fmt.Println(10.0/3) // 3.333.....

	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}

func bazz() {
	fmt.Println(user.Lookup("satoryoka"))
}
