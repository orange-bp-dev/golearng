package main

import "fmt"

func main() {
	// ()を使えば複数代入
	var (
		// = X みたく値を宣言しない場合、初期値が出力
		x int
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t, f bool = true, false
	)
	fmt.Println(x, i, f64, s, t, f)

	//型を定義しない書き方は関数内でのみ宣言可能
	//foo関数の中で変数を定義して、main関数内で実行とかとか
		xi := 1 
  //再宣言
		xi = 2
		xf64 := 1.2
		xs := "test"
		xt, xf := true, false
		fmt.Println(xi, xf64, xs, xt, xf)

		//型を調べる
		fmt.Printf("%T\n", xf64)

		// \n：改行
		//Printlnは勝手に改行されるんだけど、Printfは改行されない
		fmt.Printf("%T\n", xi)

}