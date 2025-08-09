package user

import "context"

// Repository — порт (интерфейс) для работы с пользователями.
// Его реализуют адаптеры (например, PostgreSQL).
type Repository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
