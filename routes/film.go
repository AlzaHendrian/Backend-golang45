package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoutes(e *echo.Group) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	e.GET("/films", h.FindFilm)
	e.GET("/film/:id", h.GetFilm)
	e.POST("/film", h.CreateFilm)
	e.PATCH("/film/:id", h.UpdateFilm)
	e.DELETE("/film/:id", h.DeleteFilm)
}
