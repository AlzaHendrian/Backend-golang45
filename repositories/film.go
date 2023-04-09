package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film) (models.Film, error)
	DeleteFilm(film models.Film, ID int) (models.Film, error)
	GetCategoryFilm(ID int) (models.Category, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error // add this code

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, ID).Error // add this code

	return film, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repository) UpdateFilm(film models.Film) (models.Film, error) {
	// err := r.db.Save(&film).Error
	err := r.db.Model(&film).Updates(film).Error

	return film, err
}

//

func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Delete(&film, ID).Scan(&film).Error

	return film, err
}

func (r *repository) GetCategoryFilm(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error // add this code

	return category, err
}
