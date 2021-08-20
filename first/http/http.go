package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	// //2つめはエラーハンドリング用変数
	// resp, _ := http.Get("http://example.com")
	// //これがなくてもとくに挙動変わらず
	// defer resp.Body.Close()
	// //ioutilはファイルとかレスポンスを読み込むのによく使われるらしい
	// body, _ := ioutil.ReadAll(resp.Body)
	// //バイト型で返ってくるからString型にパース
	// fmt.Println(string(body))

	//正しいurlかどうかを判定.正常であればURLが出力される
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	//base, referenceをくっつける
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	//GETの場合は第三引数はnilでOK
	//この時点ではまだリクエストを投げていない
	req, _ := http.NewRequest("GET", endpoint, nil)
	//  req.Header("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	fmt.Println(q)
}
