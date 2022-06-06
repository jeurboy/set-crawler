package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetDataFromURL(urlStocks string) (body []byte) {
	resp, err := http.Get(urlStocks)
	if err != nil {
		// handle error
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	return
}

func GetJsonFromURL(urlStocks string, target interface{}) error {
	r, err := http.Get(urlStocks)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
