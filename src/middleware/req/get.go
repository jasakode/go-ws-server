package req

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Headers struct {
	Key   string
	Value string
}

var FetchApi = func(method string, url string, headers []Headers, bodyRequest ...[]byte) (string, error) {
	// Membuat permintaan HTTP sesuai dengan metode yang diberikan
	bodyReq := bytes.NewBuffer([]byte{})
	if len(bodyRequest) > 0 {
		bodyReq = bytes.NewBuffer(bodyRequest[0])
	}
	req, err := http.NewRequest(method, url, bodyReq)
	if err != nil {
		return "", err
	}
	for _, headerValue := range headers {
		req.Header.Set(headerValue.Key, headerValue.Value)
	}

	// Melakukan permintaan HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Membaca isi respons
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
