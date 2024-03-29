package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(10, 20)
	foo(10, 20, 30)

	s:= []int{1, 2, 3}
	fmt.Println(s)
	foo(s...) //sのまま代入するとfoo({1, 2, 3})の状態なので展開してfoo(1, 2, 3)として代入する必要がある
}