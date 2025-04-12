package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/404th/smtest/internal/repository/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) Register(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	if err = a.db.Create(&req).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return resp, errors.New("username already exists")
		}
		return resp, errors.New("failed to create user: " + err.Error())
	}

	resp = req

	return
}

func (a *authRepository) Login(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	var user model.User

	data := a.db.First(&user, "username = ?", req.Username)
	if data.RowsAffected < 1 {
		err = errors.New("not found")
		return
	}

	resp = &user

	return
}

func (a *authRepository) GetUser(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	return
}
