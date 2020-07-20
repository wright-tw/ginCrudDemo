package until

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func CallApi(method string, url string, params string, header map[string]string) (string, error) {

	request, error := http.NewRequest("POST", url, strings.NewReader(params))
	if error != nil {
		return "", error
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	response, error := client.Do(request)
	defer response.Body.Close()

	if error != nil {
		return "", error
	}

	bodyByte, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyByte)

	return bodyString, nil
}
