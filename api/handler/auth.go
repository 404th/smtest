package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/404th/smtest/api/handler/validation"
	"github.com/404th/smtest/internal/repository/model"
	"github.com/404th/smtest/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	isValid, message := validation.ValidateRegister(&req)
	if !isValid {
		err := errors.New(message)
		handleErrorResponse(c, err)
		return
	}

	password, err := pkg.HashPassword(req.Password)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}
	req.Password = password

	ctx := c.Request.Context()

	resp, err := h.service.Auth().Register(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	isValid, message := validation.ValidateRegister(&req)
	if !isValid {
		err := errors.New(message)
		handleErrorResponse(c, err)
		return
	}

	ctx := c.Request.Context()

	user, err := h.service.Auth().Login(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	accessToken, err := token.SignedString([]byte(h.cfg.PasswordSalt))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
