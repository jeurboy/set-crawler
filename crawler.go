package crawler

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	httphelpers "github.com/jeurboy/set-crawler/helpers/http"
	"github.com/thoas/go-funk"

	"github.com/jeurboy/set-crawler/entity"
)

var URLScheme = "https://www.set.or.th/set/historicaltrading.do?symbol=%s&ssoPageId=2&language=th&country=TH&page=%s"
var StocklistURL = "https://www.set.or.th/dat/eod/listedcompany/static/listedCompanies_th_TH.xls"
var URLCompanyFinancial = "https://www.settrade.com/C04_03_stock_companyhighlight_p1.jsp?txtSymbol=%s&ssoPageId=12&selectPage=3"

func GetSetPriceData(stockName string, page int) (entity.PricePage, error) {
	url := fmt.Sprintf(URLScheme, stockName, strconv.Itoa(page))

	// fmt.Printf("Get data from : %s \n", url)
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

func GetAllStock() (entity.StockList, error) {
	rawPageHtml := string(httphelpers.GetDataFromURL(StocklistURL))

	//New default config
	p := entity.NewStock()

	var data entity.StockList

	//parse html data
	_ = p.Parse(&data, rawPageHtml)

	//print data
	return data, nil
}

func GetCompanyFinancial(stockName string) ([]entity.FinancialRaw, error) {
	url := fmt.Sprintf(URLCompanyFinancial, stockName)

	rawPageHtml := string(httphelpers.GetDataFromURL(url))

	//New default config
	p := entity.NewPagser()

	var data entity.FinancialPage

	//parse html data
	err := p.Parse(&data, rawPageHtml)

	//check error
	_ = err

	// Filp reo in to column
	i := 0
	if ret, ok := funk.Map(data.FinancialTable.Title, func(t string) (r entity.FinancialRaw) {
		if t == "" {
			i++
			return
		}

		if len(data.FinancialTable.FinancialData) < 9 {
			return
		}

		lineData := data.FinancialTable.FinancialData
		// if i != 0 {
		// 	funk.ForEach(lineData, func(d entity.FinancialData) {
		// 		legend := tis620.ToUTF8([]byte(d.Column[0]))
		// 		fmt.Printf("%s : %s \n", legend, d.Column[i])
		// 	})

		// }

		re := regexp.MustCompile(`\d{2}/\d{2}/\d{2}`)
		date := string(re.Find([]byte(t)))

		layout := "02/01/06"
		tp, _ := time.Parse(layout, string(date))
		tp = tp.AddDate(543, 0, 0)

		r.Date = entity.DateString(tp.Format("02/01/2006"))
		r.Asset = entity.DecimalString(lineData[0].Column[i])
		r.Liabilities = entity.DecimalString(lineData[1].Column[i])
		r.Equity = entity.DecimalString(lineData[3].Column[i])
		r.PaidUpCapital = entity.DecimalString(lineData[4].Column[i])
		r.NetProfitOrLoss = entity.DecimalString(lineData[5].Column[i])
		r.ROA = entity.DecimalString(lineData[6].Column[i])
		r.ROE = entity.DecimalString(lineData[7].Column[i])
		r.NetProfitMargin = entity.DecimalString(lineData[8].Column[i])

		i++
		return
	}).([]entity.FinancialRaw); ok {
		ret = funk.Filter(ret, func(r entity.FinancialRaw) bool {
			return r.Date != ""
		}).([]entity.FinancialRaw)
		return ret, nil
	}

	return nil, errors.New("no data")
}
