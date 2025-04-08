package repository

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"
	"gorm.io/gorm"

	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
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

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID: %s", err.Error())

		return nil, &_itemShopException.ItemNotFound{}
	}

	return item, nil
}

func (r *itemShopRepositoryImpl) buildItemFilterQuery(itemFilter *_itemShopModel.ItemFilter) *gorm.DB {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if name := strings.TrimSpace(itemFilter.Name); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if desc := strings.TrimSpace(itemFilter.Description); desc != "" {
		query = query.Where("description ILIKE ?", "%"+desc+"%")
	}

	return query
}
