package main

import "fmt"

func main() {
	// Q1. 以下の実行結果をコードを書かずに答えてください。
	// エラーがおきますか？正しく表示されるとすると何が表示字されますか？

    var i int = 10

		//ポインタの変数を定義
    var p *int

    // 変数iのアドレスを格納
    p = &i

    var j int
		//pのアドレス内に存在する変数を格納
    j = *p

		fmt.Println(j)
}