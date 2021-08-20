package main

import "fmt"

const Pi = 3.14

//constの場合は型宣言無しでOK
//実行されるタイミングで方がキマる
const (
	Username = "test_user"
	Password = "test_password"
)

func main() {
	fmt.Println(Pi, Username, Password)
}