# PapaGo
[![Build Status](https://travis-ci.com/kkweon/papago.svg?branch=master)](https://travis-ci.com/kkweon/papago)
[![GoDoc](https://godoc.org/github.com/kkweon/papago/pkg/papago?status.svg)](https://godoc.org/github.com/kkweon/papago/pkg/papago)

Naver Papago GO API Wrapper

```go
import "github.com/kkweon/papago/pkg/papago"

payload := papago.Payload{
    Source: papago.En,
    Target: papago.Ko,
    Text:   "Hello World",
}
resp, _ := papago.TranslateNMT(auth, payload)
fmt.Println(resp)
// Output: 안녕 세계
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
