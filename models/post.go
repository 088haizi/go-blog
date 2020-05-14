package models

import (
	"fmt"

	"github.com/088haizi/go-blog/database"
	"github.com/088haizi/go-blog/config"
)

func All(posts *[]database.Post) (err error) {
	if err = config.DBConn().Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func Create(post *database.Post) (err error) {
	if err = config.DBConn().Create(post).Error; err != nil {
		return err
	}
	return nil
}

func Find(post *database.Post, id string) (err error) {
	if err = config.DBConn().Where("id = ?", id).First(post).Error; err != nil {
		return err
	}
	return nil
}

func Update(post *database.Post, id string) (err error) {
	fmt.Println(post)
	config.DBConn().Save(post)
	return nil
}

func Delete(post *database.Post, id string) (err error) {
	config.DBConn().Where("id = ?", id).Delete(post)
	return nil
}
