package trade_log

import (
	"time"
	"trade_log_backend/internal/domain/currency"
)

const (
	DirectionLong  = "long"
	DirectionShort = "shorts"
)

type TradeLog struct {
	ID        int64
	UserID    int64
	Currency  currency.Currency
	Direction string
	CreatedAt time.Time
	UpdatedAt time.Time
}
