package papago

import (
	"errors"
	"os"
)

// NaverAuth contains ClientId and ClientSecret
type NaverAuth struct {
	ClientID     string
	ClientSecret string
}

// GetAuthFromEnv retrieves NaverAuth from Environment variables
//
// Set environment variables
//   #!/usr/bin/env bash
//
//   export NAVER_CLIENT_ID = "..."
//   export NAVER_CLIENT_SECRET = "..."
func GetAuthFromEnv() (NaverAuth, error) {
	clientID, ok := os.LookupEnv("NAVER_CLIENT_ID")
	if !ok {
		return NaverAuth{}, errors.New("NAVER_CLIENT_ID is not found")
	}

	clientSecret, ok := os.LookupEnv("NAVER_CLIENT_SECRET")
	if !ok {
		return NaverAuth{}, errors.New("NAVER_CLIENT_SECRET is not found")
	}

	return NaverAuth{clientID, clientSecret}, nil
}
