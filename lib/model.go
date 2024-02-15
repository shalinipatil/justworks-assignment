package lib

type CryptoResponse struct {
	Data *Data `json:"data"`
}

type Data struct {
	Curreny string            `json:"currency"`
	Rates   map[string]string `json:"rates"`
}
