package main

import (
	"fmt"

	crawler "github.com/jeurboy/set-crawler"
	"github.com/jeurboy/set-crawler/helpers/dump"
)

var URLScheme = "https://www.set.or.th/set/historicaltrading.do?symbol=%s&ssoPageId=2&language=th&country=TH&page=%s"
var StocklistURL = "https://www.set.or.th/dat/eod/listedcompany/static/listedCompanies_th_TH.xls"

func main() {
	fmt.Println("Start service")
	data, _ := crawler.GetAllStock()

	dump.DD(data)
	fmt.Println("End")
}
