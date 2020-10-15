package test

import (
	"io"
	"net/http"
	"os"
)

var env = os.Getenv("ENV")

func GetRequest(targetURL string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", targetURL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return res, nil
}

func PostRequest(targetURL string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", targetURL, body)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

func ConfigURL() string {
	if env == "local" {
		return "http://localhost:8082"
	}
	return ""
}
