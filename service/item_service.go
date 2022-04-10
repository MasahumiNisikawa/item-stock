package item

import (
	"item-stock/db"
	"item-stock/entity"

	"github.com/gin-gonic/gin"
)

type Service struct {}

type Item entity.Item

func (s Service) GetAll() ([]Item, error) {
	db := db.GetDB()
	var i []Item

	if err := db.Find(&i).Error; err != nil {
		return nil, err
	}

	return i, nil
}

func (s Service) CreateModel(c *gin.Context) (Item, error) {
	db := db.GetDB()
	var i Item

	if err := c.BindJSON(&i); err != nil {
		return i, err
	}

	if err := db.Create(&i).Error; err != nil {
		return i, err
	}
	return i, nil
}

func (s Service) GetByID(id string) (Item, error) {
	db := db.GetDB()
	var i Item

	if err := db.Where("id = ?", id).Find(&i).Error; err != nil {
		return i, err
	}
	return i, nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (Item, error) {
	db := db.GetDB()
	var i Item

	if err := db.Where("id = ?", id).Find(&i).Error; err != nil {
		return i, err
	}

	if err := c.BindJSON(&i); err != nil {
		return i, err
	}

	db.Save(&i)

	return i, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var i Item

	if err := db.Where("id = ?", id).Delete(&i).Error; err != nil {
		return err
	}
	return nil
}