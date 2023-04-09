package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	FullName  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"-" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" default:"false"`
}

type UsersProfileResponse struct {
	FullName  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"-" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" default:"false"`
}
