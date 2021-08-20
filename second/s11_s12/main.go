package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("ABC"[0]) // 65
	fmt.Println(string("ABC"[0])) // A
	s := "XXX"
	// s[0] = "Y" この書き方は出来ない！！
	// こうする！
	s = strings.Replace(s, "X", "Y", 1)
	strings.Contains(s, "Y") // 含まれますか？ => true

	x := "14"
	//int(x)は無理！ アスキーToインテジャーで 文字列 -> 整数
	i, _ := strconv.Atoi(x) 
	fmt.Println(i)


	

	
}