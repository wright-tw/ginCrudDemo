package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewClient() *Client {
	return &Client{}
}

type Client struct{}

func (c *Client) CallAPI(method string, url string, params string, header map[string]string) (string, error) {
	ctx := context.Background()
	request, error := http.NewRequest(method, url, strings.NewReader(params))
	request = request.WithContext(ctx)
	if error != nil {
		return "", error
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return "", error
	}
	defer response.Body.Close()

	bodyByte, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyByte)

	return bodyString, nil
}
