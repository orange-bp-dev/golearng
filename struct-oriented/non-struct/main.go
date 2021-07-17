package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	//直上で指定しているように返り値はintであって、MyIntではない！！
	//intに変換したくなければ、返り値の型指定の部分でMyIntを指定してあげる必要がある。
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}