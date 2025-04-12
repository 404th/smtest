package repository

import (
	"context"

	"github.com/404th/smtest/internal/repository/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) Register(ctx *context.Context, req *model.RegisterRequest) (resp *model.RegisterResponse, err error) {
	resp = &model.RegisterResponse{}

	return
}

func (a *authRepository) Login(ctx *context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error) {
	resp = &model.LoginResponse{}

	return
}

func (a *authRepository) GetUser(ctx *context.Context, req *model.GetUserRequest) (resp *model.User, err error) {
	resp = &model.User{}

	return
}
