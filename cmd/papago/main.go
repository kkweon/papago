package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkweon/papago/pkg/papago"
)

func main() {
	auth, err := papago.GetAuthFromEnv()

	if err != nil {
		panic(err)
	}

	if len(os.Args) <= 1 {
		panic("<sentence> is not given")
	}

	target := strings.Join(os.Args[1:], "")
	payload := papago.Payload{
		Source: papago.En,
		Target: papago.Ko,
		Text:   target,
	}
	resp, err := papago.TranslateNMT(auth, payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
