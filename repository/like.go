package repository

import (
	"encoding/json"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"gorm.io/gorm"
)

func InsertLike(data model.Like) (*model.Like, error) {
	db := config.DbManager()
	res := model.Like{}

	like := db.Create(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	resByte, err := json.Marshal(like.Value)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while marshal like data. Err: %s", err))
		return nil, err
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while unmarshal like data. Err: %s", err))
		return nil, err
	}

	return &res, nil
}

func GetLike(data model.Like) (*model.Like, error) {
	db := config.DbManager()
	res := model.Like{}

	if data.ArticleID > 0 {
		db = db.Where("article_id = ?", data.ArticleID)
	}

	if data.UserID > 0 {
		db = db.Where("user_id = ?", data.UserID)
	}

	if err := db.First(&res).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}

		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	return &res, nil
}

func SelectLike(data model.Like) ([]*model.Like, error) {
	db := config.DbManager()
	res := []*model.Like{}

	if data.ArticleID > 0 {
		db = db.Where("article_id = ?", data.ArticleID)
	}

	db = db.Order("id ASC")

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
