package papago

import "fmt"

func ExampleTranslateNMT() {
	auth, _ := GetAuthFromEnv()
	payload := Payload{
		Source: En,
		Target: Ko,
		Text:   "Hello World",
	}
	resp, _ := TranslateNMT(auth, payload)
	fmt.Println(resp)
	// Output: 안녕 세계
}

func ExampleTranslateSMT() {
	auth, _ := GetAuthFromEnv()
	payload := Payload{
		Source: En,
		Target: Ko,
		Text:   "Hello from Papago",
	}
	resp, _ := TranslateSMT(auth, payload)
	fmt.Println(resp)
	// Output: 안녕하세요 파파에서
}
