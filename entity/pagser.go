package entity

import (
	"time"

	"github.com/foolin/pagser"
	"github.com/foolin/pagser/extensions/markdown"
)

type PageDataRaw struct {
	Title      string `pagser:"title"`
	PriceTable []struct {
		Title     DateHeaderRaw  `pagser:"thead tr th"`
		DatePrice []DatePriceRaw `pagser:"tbody tr"`
	} `pagser:"table"`
}

//" วันที่\n    ราคาเปิด\n    ราคาสูงสุด\n    ราคาต่ำสุด\n    ราคาปิด\n    เปลี่ยนแปลง\n    %เปลี่ยนแปลง\n    ปริมาณรวม(หุ้น)\n    มูลค่ารวม('000 บาท)",
type DatePriceRaw struct {
	Date       DateString      `pagser:"td->eq(0)"`
	Open       DecimalString   `pagser:"td->eq(1)"`
	High       DecimalString   `pagser:"td->eq(2)"`
	Low        DecimalString   `pagser:"td->eq(3)"`
	Change     ChangeIntString `pagser:"td->eq(4)"`
	Volume     IntString       `pagser:"td->eq(5)"`
	TotalTrade DecimalString   `pagser:"td->eq(6)"`
}

type DateHeaderRaw []string
type DateString string
type DecimalString string
type IntString string
type ChangeIntString string

func (d DateString) ToDate() time.Time {
	return time.Now()
}

func NewPagser() *pagser.Pagser {
	p := pagser.New()
	markdown.Register(p)

	return p
}
