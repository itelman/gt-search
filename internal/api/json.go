package api

import (
	"io/ioutil"
	"net/http"
)

func ParseJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%#v\n", string(body))

	return body, nil
}
