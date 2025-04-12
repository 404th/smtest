package handler

import (
	"net/http"

	"github.com/404th/smtest/api/handler/validation"
	"github.com/404th/smtest/internal/repository/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	var req model.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, message := validation.ValidateMovie(&req)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().CreateMovie(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert a movie"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetMovieById(c *gin.Context) {
	var req model.Id
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().GetMovieById(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get a movie"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetAllMovies(c *gin.Context) {
	var req model.GetAllMoviesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().GetAllMovies(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get movies"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) DeleteMovie(c *gin.Context) {
	var req model.Id
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.service.Movie().DeleteMovie(&ctx, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get a movie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func (h *Handler) UpdateMovies(c *gin.Context) {
	var req model.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().UpdateMovie(&ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get movies"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
