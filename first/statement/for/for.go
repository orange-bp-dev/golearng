package main

import "fmt"

func main() {
	for i:=0; i<10; i++ {
		
    //if文の書き方①
		if i == 3 {
			fmt.Println("continue")
			continue //処理を終えて次の数に映る
		}

		if i == 5 {
			fmt.Println("break")
			break //以降の処理を止めてfor文を抜ける
		}
		fmt.Println(i)

		//if文の書き方②
		sum := 1

		//セミコロンは省略可能！
		for ; sum < 10; {
			sum += sum
			fmt.Println(sum)
		}
		fmt.Println(sum)
	}

	//forの条件を書かないと無限ループになる
	//for {
		//fmt.Println("Hello")
//}



}