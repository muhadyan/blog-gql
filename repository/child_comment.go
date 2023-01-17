package repository

import (
	"encoding/json"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"gorm.io/gorm"
)

func InsertChildComment(data model.ChildComment) (*model.ChildComment, error) {
	db := config.DbManager()
	res := model.ChildComment{}

	comment := db.Create(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return nil, err
	}

	resByte, err := json.Marshal(comment.Value)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while marshal child comment data. Err: %s", err))
		return nil, err
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while unmarshal child comment data. Err: %s", err))
		return nil, err
	}

	return &res, nil
}

func UpdateChildComment(data model.ChildComment) error {
	db := config.DbManager()
	childComment := model.ChildComment{}

	db = db.Model(&childComment).Where("id = ?", data.ID).Update(&data)
	err := db.Error
	if err != nil {
		err = fmt.Errorf("error while querying db. Err: %s", err)
		return err
	}

	return nil
}

func GetChildComment(data model.ChildComment) (*model.ChildComment, error) {
	db := config.DbManager()
	res := model.ChildComment{}

	if data.ID > 0 {
		db = db.Where("id = ?", data.ID)
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

func GetChildComments(data model.ChildComment) ([]*model.ChildComment, error) {
	db := config.DbManager()
	res := []*model.ChildComment{}

	if data.CommentID > 0 {
		db = db.Where("comment_id = ?", data.CommentID)
	}

	if data.IsApproved {
		db = db.Where("is_approved = ?", data.IsApproved)
	}

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
