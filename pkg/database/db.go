package database

import (
	"MyGram/app/entity"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	Config := "root:@tcp(127.0.0.1:3306)/mygram?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(Config), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.Comment{}, entity.Social{})

	return db, nil
}
