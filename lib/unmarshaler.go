package lib

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RespUnmarshaler(resp *http.Response) *CryptoResponse {
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		log.Fatal("Can Not Read Body")
	}
	obj := CryptoResponse{}

	jsonErr := json.Unmarshal(body, &obj)
	if jsonErr != nil {
		log.Fatal("Unmarshaling failed")
	}

	if obj.Data == nil || len(obj.Data.Rates) == 0 {
		log.Fatal("no data available")
	}
	return &obj
}
