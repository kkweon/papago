package papago

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const papagoNMTAPI = "https://openapi.naver.com/v1/papago/n2mt"
const papagoSMTAPI = "https://openapi.naver.com/v1/language/translate"

// TranslateNMT returns the translation by Naver NMT
func TranslateNMT(auth NaverAuth, payload Payload) (string, error) {
	return sendRequest(papagoNMTAPI, auth, payload)
}

// TranslateSMT returns the translation by Naver SMT
func TranslateSMT(auth NaverAuth, payload Payload) (string, error) {
	return sendRequest(papagoSMTAPI, auth, payload)
}

func sendRequest(url string, auth NaverAuth, payload Payload) (string, error) {
	body := payload.toFormString()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln("Unable to create a POST request")
	}
	req.Header.Add("X-Naver-Client-Id", auth.ClientID)
	req.Header.Add("X-Naver-Client-Secret", auth.ClientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var papagoResponse pResponse
	err = json.Unmarshal(respBody, &papagoResponse)

	if err != nil {
		return "", err
	}
	return papagoResponse.Message.Result.Text, nil
}
