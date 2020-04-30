package main

import (
	"fmt"
	"qoala/shortener"
)

func main() {
	testcases := []struct {
		url string
	}{
		{"https://grabkios.qoala.app?agentId=908123123&phoneNumber=085716028129"},
		{"https://offline.qoalaplus.app?agentId=908123123&phoneNumber=085716028129"},
		{"https://integration.qoala.my?agentId=908123123&phoneNumber=085716028129"},
	}
	for _, testcase := range testcases {
		encode := shortener.Encode(testcase.url)
		decode := shortener.Decode(encode)
		fmt.Printf("url: %s\n", testcase.url)
		fmt.Printf("encode: %s\n", encode)
		fmt.Printf("decode: %s\n", decode)
	}
}
