package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type GormDatabase interface {
	DbManager() *gorm.DB
}

func InitDB() (*gorm.DB, error) {
	configuration := GetConfig()
	connect_string := fmt.Sprintf("postgres://%s:%s@%s/%s", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_NAME)
	db, err = gorm.Open("postgres", connect_string)
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate(
	// 	&model.User{},
	// 	&model.Article{},
	// 	&model.Comment{},
	// 	&model.ChildComment{},
	// 	&model.Like{},
	// )

	return db, nil
}

func DbManager() *gorm.DB {
	return db
}
