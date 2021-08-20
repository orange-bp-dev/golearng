package main

import (
	"encoding/json"
	"fmt"
)

//最右の要素で、MarshalにしたときのKeyを定義できる。
//デフォルトではKeyはName, Age, Nicknamesとなる。
type Person struct {
	Name      string   `json:"-"`             //-: Marshalしたときに隠して表示しない
	Age       int      `json:"age,omitempty"` //omitempty: 値が0もしくは空文字 '', から配列 [] のときMarshalしたときに表示しない
	Nicknames []string `json:"nicknames"`     //
}

func main() {
	b := []byte(`{"name": "mike", "age": 0, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)

	//vはバイト型。String型に変換
	fmt.Println(string(v))
}
