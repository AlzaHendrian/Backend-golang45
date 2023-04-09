package filmsdto

type CreateCategoryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}
