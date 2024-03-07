package main

import (
	"encoding/base64"
	"log"
)

func main() {
	str := "Hello dudu"
	strGZIPEn := GZIPEn(str)
	log.Println(base64.StdEncoding.EncodeToString(strGZIPEn)) //加密 H4sIAAAAAAAA//JIzcnJV0gpTSkFAAAA//8BAAD//9gLokwKAAAA
	strGZIPDe, _ := GZIPDe(strGZIPEn)
	log.Println(string(strGZIPDe)) //解密

}
