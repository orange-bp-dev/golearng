// Q1. 以下に、7と表示されるメソッドを作成してください。

package main

import (
	"fmt"
)
 
type Vertex struct{
    X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func Plus(v Vertex) int {
	return v.X + v.Y
}
 
func main(){
    v := Vertex{3, 4}
    fmt.Println(v.Plus())
		fmt.Println(Plus(v))
}
