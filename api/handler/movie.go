package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/404th/smtest/api/handler/validation"
	"github.com/404th/smtest/internal/repository/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	var req model.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	isValid, message := validation.ValidateMovie(&req)
	if !isValid {
		err := errors.New(message)
		handleErrorResponse(c, err)
		return
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().CreateMovie(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetMovieById(c *gin.Context) {
	var req model.Id
	if c.Param("id") == "" {
		err := errors.New("id kiritilishi kerak")
		handleErrorResponse(c, err)
		return
	}

	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := errors.New("id noto'g'ri formatda kiritilgan")
		handleErrorResponse(c, err)
		return
	}

	req.Id = uint(idInt)

	ctx := c.Request.Context()

	resp, err := h.service.Movie().GetMovieById(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetAllMovies(c *gin.Context) {
	var (
		req      model.GetAllMoviesRequest
		metadata model.Metadata
	)

	limitInt, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		err := errors.New("limit noto'g'ri kiritilgan")
		handleErrorResponse(c, err)
		return
	}

	pageInt, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		err := errors.New("page noto'g'ri kiritilgan")
		handleErrorResponse(c, err)
		return
	}

	metadata.Limit = uint(limitInt)
	metadata.Page = uint(pageInt)
	req.Metadata = &metadata

	req.Search = c.Query("search")

	if metadata.Limit < 10 {
		metadata.Limit = 10
	} else if metadata.Limit > 100 {
		metadata.Limit = 100
	}

	if metadata.Page < 1 {
		metadata.Page = 1
	}

	ctx := c.Request.Context()

	resp, err := h.service.Movie().GetAllMovies(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) DeleteMovie(c *gin.Context) {
	var req model.Id

	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := errors.New("id noto'g'ri formatda kiritilgan")
		handleErrorResponse(c, err)
		return
	}

	req.Id = uint(idInt)

	ctx := c.Request.Context()

	if err := h.service.Movie().DeleteMovie(&ctx, &req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Muvaffaqiyatli o'chirildi"})
}

func (h *Handler) UpdateMovies(c *gin.Context) {
	var req model.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, err)
		return
	}

	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := errors.New("id noto'g'ri formatda kiritilgan")
		handleErrorResponse(c, err)
		return
	}

	req.Id = uint(idInt)

	ctx := c.Request.Context()

	resp, err := h.service.Movie().UpdateMovie(&ctx, &req)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
