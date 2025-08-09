package db

import (
	"context"
	"time"

	"trade_log_backend/internal/domain/currency"
	"trade_log_backend/internal/domain/trade_log"

	"gorm.io/gorm"
)

// TradeLogModel — структура для работы с БД через GORM
type TradeLogModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"not null;index"`
	Currency  string    `gorm:"size:10;not null"` // будем хранить код валюты, например "USD"
	Direction string    `gorm:"size:10;not null"` // например "buy" или "sell"
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func fromDomain(tl *trade_log.TradeLog) *TradeLogModel {
	return &TradeLogModel{
		ID:        tl.ID,
		UserID:    tl.UserID,
		Currency:  tl.Currency.Code,
		Direction: tl.Direction,
		CreatedAt: tl.CreatedAt,
		UpdatedAt: tl.UpdatedAt,
	}
}

func toDomain(m *TradeLogModel) (*trade_log.TradeLog, error) {
	currency := &currency.Currency{}
	currency, err := currency.GetCurrencyByCode(m.Currency)
	if err != nil {
		return nil, err
	}
	return &trade_log.TradeLog{
		ID:        m.ID,
		UserID:    m.UserID,
		Currency:  *currency,
		Direction: m.Direction,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

// TradeLogRepository — репозиторий для работы с trade_log
type TradeLogRepository struct {
	db *gorm.DB
}

func NewTradeLogRepository(db *gorm.DB) *TradeLogRepository {
	return &TradeLogRepository{db: db}
}

func (r *TradeLogRepository) GetByID(ctx context.Context, id int64) (*trade_log.TradeLog, error) {
	var m TradeLogModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return toDomain(&m), nil
}

func (r *TradeLogRepository) Create(ctx context.Context, tl *trade_log.TradeLog) error {
	m := fromDomain(tl)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	tl.ID = m.ID
	return nil
}

func (r *TradeLogRepository) Update(ctx context.Context, tl *trade_log.TradeLog) error {
	m := fromDomain(tl)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *TradeLogRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&TradeLogModel{}, id).Error
}
