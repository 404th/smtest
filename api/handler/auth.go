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

// Register
// @ID			register
// @Router		/register [POST]
// @Summary		Register Client
// @Description	Register Client
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		object	body		model.RegisterRequest	true	"RegisterRequest"
// @Success		200		{object}	model.Response{date=model.User}	"Response"
// @Response	400		{object}	model.Response{}				"Invalid Argument"
// @Failure		500		{object}	model.Response{}				"Server Error"
func (h *Handler) Register(c *gin.Context) {
	var req model.RegisterRequest
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

	resp, err := h.service.Auth().Register(&ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Message:    "OK",
		StatusCode: http.StatusOK,
		Data:       resp,
	})
}

// Login
// @ID			login
// @Router		/login [POST]
// @Summary		Login Client
// @Description	Login Client
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		object	body		model.LoginRequest				true			"User"
// @Success		200		{object}	model.Response{data=model.User}					"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	isValid, message := validation.ValidateRegister(&model.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if !isValid {
		err := errors.New(message)
		handleErrorResponse(c, err)
		return
	}

	ctx := c.Request.Context()

	user, err := h.service.Auth().Login(&ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	accessToken, err := token.SignedString([]byte(h.cfg.JWTSecretSalt))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
