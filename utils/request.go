package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(host string, postVaules url.Values) (body []byte, err error) {
	// uri := url.URL{
	// 	Host: host,
	// }
	postString := postVaules.Encode()

	req, err := http.NewRequest("POST", host, strings.NewReader(postString))

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return
}
