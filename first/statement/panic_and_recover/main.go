package main

import "fmt"

func thirdPartyConnectDB() {
	//panic: 強制終了
	panic("Unable to connect database !")
}

func save() {
	// defer: 関数の最後に実行
	//仮にthirdPartyConnectDB()をdeferの前に置くとrecover
	defer func(){
		//recover: panicで強制終了した処理を再開
		s := recover()
		fmt.Println(s)
	}()

	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}