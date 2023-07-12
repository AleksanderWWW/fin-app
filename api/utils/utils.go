package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)


func createRequest(params map[string]string, headers map[string]string, baseUrl string) (*http.Request, error) {
	parameters := url.Values{}

	var urlStr string

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}

		u, err := url.ParseRequestURI(baseUrl)
		if err != nil {
			return nil, err
		}

		u.RawQuery = parameters.Encode()
		urlStr = fmt.Sprintf("%v", u)
	} else {
		urlStr = baseUrl
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func GetResponse(params map[string]string, headers map[string]string, baseUrl string) (string, error) {
	req, err := createRequest(params, headers, baseUrl)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respText), nil
}
