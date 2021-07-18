package main

import "fmt"

func main(){
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	//取り出したぶんだけ初期キャパの2が増える
	x := <-ch
	fmt.Println(x)

	//OK
	ch <- 300
	fmt.Println(len(ch))

	//range等でループでchannelの中身を取得するときは直前でchannelをcloseしてあげる必要がある
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}

}