package entity

import (
	"strconv"
	"strings"
	"time"

	"github.com/foolin/pagser"
	"github.com/shopspring/decimal"
)

type PricePage struct {
	Title      string `pagser:"title"`
	PriceTable struct {
		// Title     DateHeaderRaw  `pagser:"thead tr th"`
		DatePrice []DatePriceRaw `pagser:"tbody tr"`
	} `pagser:"table"`
}

//" วันที่\n    ราคาเปิด\n    ราคาสูงสุด\n    ราคาต่ำสุด\n    ราคาปิด\n    เปลี่ยนแปลง\n    %เปลี่ยนแปลง\n    ปริมาณรวม(หุ้น)\n    มูลค่ารวม('000 บาท)",
type DatePriceRaw struct {
	Date       DateString    `pagser:"td->eq(0)" json:"date"`
	Open       DecimalString `pagser:"td->eq(1)" json:"open"`
	High       DecimalString `pagser:"td->eq(2)" json:"high"`
	Low        DecimalString `pagser:"td->eq(3)" json:"low"`
	Close      DecimalString `pagser:"td->eq(5)" json:"close"`
	Change     DecimalString `pagser:"td->eq(6)" json:"change"`
	Volume     IntString     `pagser:"td->eq(8)" json:"totalVolume"`
	TotalTrade DecimalString `pagser:"td->eq(9)" json:"totalValue"`
}

type DateHeaderRaw []string
type DateString string
type DecimalString string
type IntString string
type ChangeIntString string

func (das DateString) ToDate() (t time.Time) {
	layout := "02/01/06"
	tp, err := time.Parse(layout, string(das))

	if err != nil {
		return
	}
	return tp
}

func (des DecimalString) ToDecimal() decimal.Decimal {
	str := strings.Replace(string(des), ",", "", -1)
	if ret, err := decimal.NewFromString(string(str)); err == nil {
		return ret
	}

	ret, _ := decimal.NewFromString("0.00")

	return ret
}

func (dei IntString) ToInt() int32 {
	str := strings.Replace(string(dei), ",", "", -1)

	if retint, err := strconv.Atoi(str); err == nil {
		return int32(retint)
	}

	return 0
}

func NewPagser() *pagser.Pagser {
	p := pagser.New()

	return p
}
