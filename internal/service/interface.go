package service

import (
	"context"
	"github.com/sorawaslocked/car-rental-user-service/internal/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) (uint64, error)
	FindOne(ctx context.Context, filter model.UserFilter) (model.User, error)
	Find(ctx context.Context, filter model.UserFilter) ([]model.User, error)
	Update(ctx context.Context, filter model.UserFilter, update model.UserUpdate) error
	Delete(ctx context.Context, filter model.UserFilter) error
}

type JwtProvider interface {
	GenerateAccessToken(id uint64, roles []string) (string, error)
	GenerateRefreshToken(id uint64, roles []string) (string, error)
	VerifyAndParseClaims(token string) (uint64, []string, error)
}
