package main

import (
	"api/entity"
	"api/helpers/dump"
	"fmt"
	"io/ioutil"
	"net/http"
)

var URLScheme = "https://www.set.or.th/set/historicaltrading.do?symbol=%s&ssoPageId=2&language=th&country=TH&page=%s"

func main() {
	data, _ := GetSetPriceData("AAV", 0)

	//print data
	dump.DD(data)

	fmt.Println("End")
}

func GetSetPriceData(stockName string, page int) (entity.PageDataRaw, error) {
	url := fmt.Sprintf(URLScheme, stockName, page)
	rawPageHtml := string(getDataFromURL(url))

	//New default config
	p := entity.NewPagser()

	var data entity.PageDataRaw

	//parse html data
	err := p.Parse(&data, rawPageHtml)

	//check error
	if err != nil {
		return data, err
	}

	//print data
	return data, nil
}

func getDataFromURL(urlStocks string) (body []byte) {
	resp, err := http.Get(urlStocks)
	if err != nil {
		// handle error
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	return
}
