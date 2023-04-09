package database

import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

func RunMigrations() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Film{},
		&models.Category{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
