package papago

import "testing"

func TestPapagoResponseDecode(t *testing.T) {
	response := `{"message":{"@type":"response","@service":"naverservice.nmt.proxy","@version":"1.0.0","result":{"srcLangType":"en","tarLangType":"ko","translatedText":"안녕 세 계"}}}`

	papagoResponse, err := decodePapagoResponse([]byte(response))

	if err != nil {
		t.Fatalf("Decode has failed")
	}

	expected := pResponse{
		Message: pResponseMessage{
			Type:    "response",
			Service: "naverservice.nmt.proxy",
			Version: "1.0.0",
			Result: Response{
				Source: "en",
				Target: "ko",
				Text:   "안녕 세 계",
			},
		},
	}

	if papagoResponse != expected {
		t.Fatal("Expected", expected, "but got", papagoResponse)
	}
}

func TestPapagoResponseDecodeFail(t *testing.T) {
	response := `{"wronginput":{}}`

	_, err := decodePapagoResponse([]byte(response))

	if err != nil {
		return
	}
	t.Fatal("Invalid json should fail this function")
}
