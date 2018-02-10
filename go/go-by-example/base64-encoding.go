package main

import b64 "encoding/base64"
import "fmt"

func main() {
	encode := b64.StdEncoding.EncodeToString
	decode := b64.StdEncoding.DecodeString
	puts := fmt.Println
	data := "abc123!?$*&()'-=@~"

	sEnc := encode([]byte(data))
	puts(sEnc)

	sDec, _ := decode(sEnc)
	puts(string(sDec))
	puts()

	uEnc := encode([]byte(data))
	puts(uEnc)
	uDec, _ := decode(uEnc)
	puts(string(uDec))
}
