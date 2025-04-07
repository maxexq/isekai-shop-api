package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/entities"
	"gorm.io/gorm"

	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{})

	if itemFilter.Name != "" {
		query = query.Where("name like ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description like ?", "%"+itemFilter.Description+"%")
	}

	if err := query.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())

		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
