package domain

import (
	"context"

	"gorm.io/gorm"
)

// Model Database
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// Kontrak untuk Repository (Ke Database)
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

// Kontrak untuk Usecase (Logic Bisnis)
type UserUsecase interface {
	Register(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (string, error) // Return JWT
}
