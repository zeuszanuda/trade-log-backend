package currency

import "errors"

var currencyMap map[string]string = map[string]string{
	"USDT": "Tether",
	"BTC":  "Bitcoin",
	"ETH":  "Ethereum",
	"XRP":  "XRP",
	"BNB":  "BNB",
	"SOL":  "Solana",
	"DOGE": "Dogecoin",
	"ADA":  "Cardano",
	"HYPE": "Hyperliquid",
	"TON":  "Toncoin",
	"PEPE": "Pepe",
}

type Currency struct {
	Title string
	Code  string
}

func (*Currency) GetCurrencyByCode(code string) (*Currency, error) {
	title, ok := currencyMap[code]
	if !ok {
		return nil, errors.New("unknown currency code")
	}

	return &Currency{
		Title: title,
		Code:  code,
	}, nil
}
