package itemshoprepository

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
	"gorm.io/gorm"

	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
)

type IRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error
	ListItem(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	CountItem(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindItemByID(itemID uint64) (*entities.Item, error)
	FindItemListByID(itemIDs []uint64) ([]*entities.Item, error)
	RecordPurchaseHistory(tx *gorm.DB, purchaseEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
}

type RepositoryService struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) IRepository {
	return &RepositoryService{db, logger}
}

func (r *RepositoryService) TransactionBegin() *gorm.DB {
	return r.db.Connect().Begin()
}

func (r *RepositoryService) TransactionRollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *RepositoryService) TransactionCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *RepositoryService) ListItem(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
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

func (r *RepositoryService) CountItem(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.buildItemFilterQuery(itemFilter)

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("CountItem items failed: %s", err.Error())

		return -1, &_itemShopException.ItemCounting{}
	}

	return count, nil
}

func (r *RepositoryService) FindItemByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID: %s", err.Error())

		return nil, &_itemShopException.ItemNotFound{}
	}

	return item, nil
}

func (r *RepositoryService) FindItemListByID(itemIDs []uint64) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if err := r.db.Connect().Model(&entities.Item{}).Where("id in ?", itemIDs).Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to find items by ID list: %s", err.Error())

		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}

func (r *RepositoryService) RecordPurchaseHistory(tx *gorm.DB, purchaseEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	insertedPurchaseHistory := new(entities.PurchaseHistory)

	if err := conn.Create(purchaseEntity).Scan(insertedPurchaseHistory).Error; err != nil {
		r.logger.Errorf("Creating purchase history failed: %s", err.Error())

		return nil, &_itemShopException.HistoryOfPurchaseRecording{}
	}

	return insertedPurchaseHistory, nil
}
