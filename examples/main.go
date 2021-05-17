package main

import (
	"fmt"

	crawler "github.com/jeurboy/set-crawler"
	"github.com/jeurboy/set-crawler/helpers/dump"
	// "github.com/jeurboy/set-crawler/helpers/dump"
)

func main() {
	fmt.Println("Start service")

	data, _ := crawler.GetAllStock()
	dump.DD(data.Stocks[0].Stocks)

	data2, _ := crawler.GetSetPriceData("AAV", 0)
	dump.DD(data2.PriceTable.DatePrice)

	data3, _ := crawler.GetCompanyFinancial("OCEAN")
	dump.DD(data3)

	fmt.Println("End")
}
