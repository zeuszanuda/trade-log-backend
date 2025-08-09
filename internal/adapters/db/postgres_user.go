package db

import (
	"context"
	"time"

	"gorm.io/gorm"

	"trade_log_backend/internal/domain/user"
)

type UserModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"`
	Name      string    `gorm:"size:255;not null"`
	Role      string    `gorm:"size:50;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func UserModelfromDomain(u *user.User) *UserModel {
	return &UserModel{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Role:      string(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserModeltoDomain(m *UserModel) *user.User {
	return &user.User{
		ID:        m.ID,
		Email:     m.Email,
		Name:      m.Name,
		Role:      user.Role(m.Role),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	var m UserModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return UserModeltoDomain(&m), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var m UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}
	return UserModeltoDomain(&m), nil
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	m := fromDomain(u)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	u.ID = m.ID
	return nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	m := fromDomain(u)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, id).Error
}
