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

func ParseJsonFile(dir string) ([]byte, error) {
	content, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WriteJsonFile(url, dir string) error {
	body, err := ParseJson(url)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dir, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
