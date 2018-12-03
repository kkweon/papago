package papago

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Payload represents a payload sent to the API
//
// Source/Target can be "ko" | "en" | "zh-CN" | "zh-TW" | "es" | "fr" | "vi" | "th" | "id"
//
// For more details, https://developers.naver.com/docs/nmt/reference/#3--%EC%9A%94%EC%B2%AD-%EB%B3%80%EC%88%98
type Payload struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Text   string `json:"text"`
}

// Response represents a response from the API
//
// Currently, it's same as Payload, but it may change in the future
type Response struct {
	Source string `json:"srcLangType"`
	Target string `json:"tarLangType"`
	Text   string `json:"translatedText"`
}

type pResponse struct {
	Message pResponseMessage
}

type pResponseMessage struct {
	Type    string   `json:"@type"`
	Service string   `json:"@service"`
	Version string   `json:"@version"`
	Result  Response `json:"result"`
}

// FormEncode
func (p Payload) toFormString() []byte {
	return []byte(fmt.Sprintf("source=%v&target=%v&text=%v", p.Source, p.Target, p.Text))
}

// Decode Response
func decodePapagoResponse(resp []byte) (pResponse, error) {
	var papagoResponse pResponse
	err := json.Unmarshal(resp, &papagoResponse)

	if (err != nil || papagoResponse.Message == pResponseMessage{} || papagoResponse.Message.Result == Response{}) {
		return papagoResponse, errors.New("resp was not valid for PapagoResponse")
	}

	return papagoResponse, nil
}
