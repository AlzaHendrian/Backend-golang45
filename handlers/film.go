package handlers

import (
	filmsdto "dumbmerch/dto/film"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FilmHandler struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *FilmHandler {
	return &FilmHandler{FilmRepository}
}

func (h *FilmHandler) FindFilm(c echo.Context) error {
	films, err := h.FilmRepository.FindFilm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *FilmHandler) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *FilmHandler) CreateFilm(c echo.Context) error {
	request := new(filmsdto.CreateFilmRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	film := models.Film{
		Title:         request.Title,
		ThumbnailFilm: request.ThumbnailFilm,
		Year:          request.Year,
		Category:      request.Category,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
	}

	data, err := h.FilmRepository.CreateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)})
}

func (h *FilmHandler) UpdateFilm(c echo.Context) error {
	request := new(filmsdto.UpdateFilmRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	fmt.Println(request, "ini request")

	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	fmt.Println(film, "ini film")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	film.Title = "test123"

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}

	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}

	if request.Year != 0 {
		film.Year = request.Year
	}
	if request.Description != "" {
		film.Description = request.Description
	}

	dataCategory, _ := h.FilmRepository.GetCategoryFilm(film.CategoryID)

	fmt.Println(film, "bebas")

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	data.Category = dataCategory

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)})
}

func (h *FilmHandler) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FilmRepository.DeleteFilm(film, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)})
}

func convertResponseFilm(u models.Film) filmsdto.FilmResponse {
	return filmsdto.FilmResponse{
		ID:            u.ID,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		Year:          u.Year,
		Category:      u.Category,
		CategoryID:    u.CategoryID,
		Description:   u.Description,
	}
}
