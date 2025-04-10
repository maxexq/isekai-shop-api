package repository

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"

	_itemManagingException "github.com/maxexq/isekei-shop-api/pkg/itemManaging/exception"
	_itemManagingModel "github.com/maxexq/isekei-shop-api/pkg/itemManaging/model"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{
		db,
		logger,
	}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {

		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint") {
			return nil, &_itemManagingException.ItemAlreadyExists{}
		}

		r.logger.Errorf("Creating item failed: %s", err.Error())

		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())

		return &_itemManagingException.ItemArchiving{}
	}

	return nil
}
