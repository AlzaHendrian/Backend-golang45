package models

type Film struct {
	ID            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string   `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          int      `json:"year"`
	Category      Category `json:"category"`
	CategoryID    int      `json:"category_id"`
	Description   string   `json:"description" gorm:"type: varchar(255)"`
}
