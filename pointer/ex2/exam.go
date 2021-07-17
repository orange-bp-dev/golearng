//Q2. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示されますか？

package main

import (
	"fmt"
)
 
func main() {
    var i int = 100
    var j int = 200
    var p1 *int
    var p2 *int
    p1 = &i
    p2 = &j
    i = *p1 + *p2 //アドレス同士の足し算は出来ないのでエラーが起きると予想 => おきなかった！！
    p2 = p1
    j = *p2 + i
    fmt.Println(j)
}