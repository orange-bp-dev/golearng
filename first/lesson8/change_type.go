package main

import (
	"fmt"
	"strconv"
)

func main(){
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	fmt.Printf("%T %v", xx, xx)

	var s string = "14"
	
	//ss := int(s) //これは無理

	// _ ではなくerrを使うと使用しないと実行時にエラーが出るので、使用しない変数は _ で定義
	// Atoi : Aschey to Integer
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T, %v\n", i, i)
	
}