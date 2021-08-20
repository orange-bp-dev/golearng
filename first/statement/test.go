package main

import (
	"fmt"
	"sort"
)

func main() {
	//Q1. 以下のスライスから最大値・最小値を探して出力するコードを書いてください
	//俺の回答
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
  sort.Sort(sort.IntSlice(l))
	fmt.Println(l[0]) //最小値
	fmt.Println(l[len(l)-1]) //最大値

	//模範解答
	var min int
 
    for i, num := range l {
        if i == 0 {
            min = num
            continue
        }
 
        if min >= num {
            min = num
        }
    }
 
    fmt.Println(min)

	//Q2. 以下の果物の価格の合計を出力するコードを書いてください。
	//俺の回答(模範解答と同じ！えらい！！)
	m := map[string]int{
		"apple":  200,
    "banana": 300,
    "grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
	}

	a := 0

  for _, v := range m{
		 a+=v
	}

	fmt.Println(a)
}

