package repository

import (
	"encoding/json"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"gorm.io/gorm"
)

func InsertUser(data model.User) (model.User, error) {
	db := config.DbManager()
	res := model.User{}

	user := db.Create(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return res, err
	}

	resByte, err := json.Marshal(user.Value)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while marshal user data. Err: %s", err))
		return res, err
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while unmarshal user data. Err: %s", err))
		return res, err
	}

	return res, nil
}

func GetUser(data model.User) (*model.User, error) {
	db := config.DbManager()
	user := model.User{}

	if data.ID > 0 {
		db = db.Where("id = ?", data.ID)
	}

	if data.Username != "" {
		db = db.Where("username = ?", data.Username)
	}

	if data.Email != "" {
		db = db.Where("email = ?", data.Email)
	}

	if data.Token != nil {
		db = db.Where("token = ?", data.Token)
	}

	if data.IsActive {
		db = db.Where("is_active = ?", data.IsActive)
	}

	if err := db.First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUser(data model.User) (*model.User, error) {
	db := config.DbManager()
	res := model.User{}

	db = db.Model(&res).Where("id = ?", data.ID).Update(&data).First(&res)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	return &res, nil
}
