package repository

import (
	"strings"

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

	query := r.buildItemFilterQuery(itemFilter)

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())

		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.buildItemFilterQuery(itemFilter)

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting items failed: %s", err.Error())

		return -1, &_itemShopException.ItemCounting{}
	}

	return count, nil
}

func (r *itemShopRepositoryImpl) buildItemFilterQuery(itemFilter *_itemShopModel.ItemFilter) *gorm.DB {
	query := r.db.Model(&entities.Item{}).Where("is_achive = ?", false)

	if name := strings.TrimSpace(itemFilter.Name); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if desc := strings.TrimSpace(itemFilter.Description); desc != "" {
		query = query.Where("description ILIKE ?", "%"+desc+"%")
	}

	return query
}
