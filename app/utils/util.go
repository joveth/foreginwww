package utils

import (
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

func GetHTMLContent(url string) (string, error) {
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr, err
	}
	return "", err
}
