package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	const apikey = "UserKey"
	const apiSecret = "UserSecret"

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte("data"))
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

}
