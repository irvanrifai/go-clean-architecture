package usecase

import (
	"context"

	"github.com/irvanrifai/go-clean-architecture/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(ctx context.Context, email, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &domain.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return "dummy-jwt-token", nil // Nanti kita ganti dengan JWT beneran
}
