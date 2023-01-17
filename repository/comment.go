package repository

import (
	"encoding/json"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"gorm.io/gorm"
)

func InsertComment(data model.Comment) (*model.Comment, error) {
	db := config.DbManager()
	res := model.Comment{}

	comment := db.Create(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	resByte, err := json.Marshal(comment.Value)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while marshal comment data. Err: %s", err))
		return nil, err
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while unmarshal comment data. Err: %s", err))
		return nil, err
	}

	return &res, nil
}

func GetComment(data model.Comment) (*model.Comment, error) {
	db := config.DbManager()
	res := model.Comment{}

	if data.ID > 0 {
		db = db.Where("id = ?", data.ID)
	}

	if data.IsApproved {
		db = db.Where("is_approved = ?", data.IsApproved)
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

func UpdateComment(data model.Comment) error {
	db := config.DbManager()
	comment := model.Comment{}

	db = db.Model(&comment).Where("id = ?", data.ID).Update(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return err
	}

	return nil
}

func SelectComment(request model.GetCommentsRequest) ([]*model.Comment, error) {
	db := config.DbManager()
	res := []*model.Comment{}

	if request.ArticleID > 0 {
		db = db.Where("article_id = ?", request.ArticleID)
	}

	db = db.Where("is_approved = ?", true)

	if request.OrderBy != "" && request.SortBy != "" {
		order := fmt.Sprintf(`%s %s`, request.OrderBy, request.SortBy)
		db = db.Order(order)
	} else {
		db = db.Order(`id ASC`)
	}

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
