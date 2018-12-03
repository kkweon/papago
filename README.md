# PapaGo
[![Build Status](https://travis-ci.com/kkweon/papago.svg?branch=master)](https://travis-ci.com/kkweon/papago)

Naver Papago GO API Wrapper

```go
func ExampleTranslateNMT() {
	auth, _ := GetAuthFromEnv()
	payload := Payload{
		Source: "en",
		Target: "ko",
		Text:   "Hello World",
	}
	resp, _ := TranslateNMT(auth, payload)
	fmt.Println(resp)
	// Output: 안녕 세계
}

func ExampleTranslateSMT() {
	auth, _ := GetAuthFromEnv()
	payload := Payload{
		Source: "en",
		Target: "ko",
		Text:   "Hello from Papago",
	}
	resp, _ := TranslateSMT(auth, payload)
	fmt.Println(resp)
	// Output: 안녕하세요 파파에서
}
```

## Get Started

1. Set Environment Variables or use `papago.NaverAuth{clientID, clientSecret}`

```bash
export NAVER_CLIENT_ID = "..."
export NAVER_CLIENT_SECRET = "..."
```

2. Retrieve `papago.NaverAuth`
```go
auth, err := papago.GetAuthFromEnv()
```

3. 2 Functions to use

```go
papago.TranslateNMT(auth, payload) (string, error)
papago.TranslateSMT(auth, payload) (string, error)
```
