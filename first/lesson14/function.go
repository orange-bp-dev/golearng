package main

import "fmt"

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

//引数で返り値を命名
func cal(price, item int)(result int) {
result = price * item
return
}

func main(){
 r1, r2 :=	add(10, 20)
 fmt.Println(r1, r2)

 r3 := cal(100, 2)
 fmt.Println(r3)

 f := func() {
	 fmt.Println("inner func")
 }
 f()

 //関数を定義すると同時に実行もできる！！
 func(x int) {
	 fmt.Println("inner func", x)
 }(1)
 
}