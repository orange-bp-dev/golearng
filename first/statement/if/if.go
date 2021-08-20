package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {

	//２要素で書く方法
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}

  //１要素でまとめて各方法
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great2")
	}
}