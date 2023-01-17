package repository

import (
	"encoding/json"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"gorm.io/gorm"
)

func InsertArticle(data model.Article) (*model.Article, error) {
	db := config.DbManager()
	res := model.Article{}

	article := db.Create(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	resByte, err := json.Marshal(article.Value)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while marshal article data. Err: %s", err))
		return nil, err
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while unmarshal article data. Err: %s", err))
		return nil, err
	}

	return &res, nil
}

func GetArticle(data model.Article) (*model.Article, error) {
	db := config.DbManager()
	res := model.Article{}

	if data.ID > 0 {
		db = db.Where("id = ?", data.ID)
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

func UpdateArticle(data model.Article) (*model.Article, error) {
	db := config.DbManager()
	res := model.Article{}

	db = db.Model(&res).Where("id = ?", data.ID).Update(&data).First(&res)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	return &res, nil
}

func DeleteArticle(data model.Article) error {
	db := config.DbManager()
	article := model.Article{}

	db = db.Where("id = ?", data.ID).Delete(&article)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return err
	}

	return nil
}

func SelectArticles(request model.GetArticlesRequest) ([]*model.Article, error) {
	db := config.DbManager()
	res := []*model.Article{}

	if request.Search != nil {
		db = db.Where("title ILIKE ?", "%"+*request.Search+"%")
	}

	if request.OrderBy != "" && request.SortBy != "" {
		order := fmt.Sprintf(`%s %s`, request.OrderBy, request.SortBy)
		db = db.Order(order)
	} else {
		db = db.Order(`id DESC`)
	}

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
