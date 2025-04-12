package handler

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, message := validation.ValidateRegister(&req)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	password, err := pkg.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Password = password

	ctx := c.Request.Context()

	resp, err := h.service.Auth().Register(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, message := validation.ValidateRegister(&req)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	ctx := c.Request.Context()

	user, err := h.service.Auth().Login(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	accessToken, err := token.SignedString([]byte(h.cfg.PasswordSalt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
