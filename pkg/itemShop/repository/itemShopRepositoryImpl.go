package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/entities"
	"gorm.io/gorm"

	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())

		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
