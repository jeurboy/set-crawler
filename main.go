package main

import (
	"api/entity"
	"api/helpers/dump"
	httphelpers "api/helpers/http"
	"fmt"
)

var URLScheme = "https://www.set.or.th/set/historicaltrading.do?symbol=%s&ssoPageId=2&language=th&country=TH&page=%s"
var StocklistURL = "https://www.set.or.th/dat/eod/listedcompany/static/listedCompanies_th_TH.xls"

func main() {
	data, _ := GetSetPriceData("AAV", 0)

	//print data
	dump.DD(data)

	// GetAllStock(StocklistURL)

	fmt.Println("End")
}

func GetSetPriceData(stockName string, page int) (entity.PricePage, error) {
	url := fmt.Sprintf(URLScheme, stockName, page)
	rawPageHtml := string(httphelpers.GetDataFromURL(url))

	//New default config
	p := entity.NewPagser()

	var data entity.PricePage

	//parse html data
	err := p.Parse(&data, rawPageHtml)

	//check error
	if err != nil {
		return data, err
	}

	//print data
	return data, nil
}

func GetAllStock(url string) (entity.StockList, error) {
	rawPageHtml := string(httphelpers.GetDataFromURL(url))

	//New default config
	p := entity.NewStock()

	var data entity.StockList

	//parse html data
	_ = p.Parse(&data, rawPageHtml)

	//print data
	return data, nil
}
