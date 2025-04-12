package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/404th/smtest/api/handler/validation"
	"github.com/404th/smtest/internal/repository/model"
	"github.com/gin-gonic/gin"
)

// CreateMovie
// @ID			createMovie
// @Security	ApiKeyAuth
// @Router		/movies [POST]
// @Summary		Create Movie
// @Description	Create Movie
// @Tags		movies
// @Accept		json
// @Produce		json
// @Param		object	body		model.CreateMovieRequest		true			"User"
// @Success		200		{object}	model.Response{data=model.Id}					"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
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

// GetMovieById
// @ID			getMovieById
// @Security	ApiKeyAuth
// @Router		/movies/{id} [GET]
// @Summary		Get Movie By Id
// @Description	Get Movie By Id
// @Tags		movies
// @Accept		json
// @Produce		json
// @Param		id		path		string									true	"id"
// @Success		200		{object}	model.Response{data=model.User}					"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
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

// GetAllMovies
// @ID			getAllMovies
// @Security	ApiKeyAuth
// @Router		/movies [GET]
// @Summary		Get All Movies
// @Description	Get All Movies
// @Tags		movies
// @Accept		json
// @Produce		json
// @Param		offset	query		integer									false	"page"
// @Param		limit	query		integer									false	"limit"
// @Param		search	query		string									false	"search"
// @Success		200		{object}	model.Response{data=model.GetAllMoviesResponse}	"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
func (h *Handler) GetAllMovies(c *gin.Context) {
	var (
		req      model.GetAllMoviesRequest
		metadata model.Metadata
	)

	limitInt, err := strconv.Atoi(c.Query("limit"))
	if err != nil && c.Query("limit") != "" {
		err := errors.New("limit noto'g'ri kiritilgan")
		handleErrorResponse(c, err)
		return
	} else if c.Query("limit") == "" {
		limitInt = 1
	}

	pageInt, err := strconv.Atoi(c.Query("page"))
	if err != nil && c.Query("page") != "" {
		err := errors.New("page noto'g'ri kiritilgan")
		handleErrorResponse(c, err)
		return
	} else if c.Query("page") == "" {
		pageInt = 1
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

// DeleteMovie
// @ID			deleteMovie
// @Security	ApiKeyAuth
// @Router		/movies/{id} [DELETE]
// @Summary		Delete Movie
// @Description	Delete Movie
// @Tags		movies
// @Accept		json
// @Produce		json
// @Param		id		path		string									true	"id"
// @Success		200		{object}	model.Response{data=nil}						"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
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

// UpdateMovies
// @ID			updateMovies
// @Security	ApiKeyAuth
// @Router		/movies/{id} [PUT]
// @Summary		Update Movie
// @Description	Update Movie
// @Tags		movies
// @Accept		json
// @Produce		json
// @Param		id		path		string								true		"Id"
// @Param		object	body		model.UpdateMovieRequestSwagger		true		"UpdateMovieRequest"
// @Success		200		{object}	model.Response{data=model.Movie}				"Response"
// @Response	400		{object}	model.Response{}								"Invalid Argument"
// @Response	404		{object}	model.Response{}								"Invalid Argument"
// @Failure		500		{object}	model.Response{}								"Server Error"
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
