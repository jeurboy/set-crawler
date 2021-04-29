package entity

import (
	"github.com/foolin/pagser"
)

type StockList struct {
	Stocks []struct {
		Title  DateHeaderRaw `pagser:"thead tr th"`
		Stocks []Stock       `pagser:"tbody tr"`
	} `pagser:"table"`
}

type Stock struct {
	Name string `pagser:"td->eq(0)"`
}

func NewStock() *pagser.Pagser {
	p := pagser.New()

	return p
}
