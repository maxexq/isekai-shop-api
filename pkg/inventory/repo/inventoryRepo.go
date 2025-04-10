package repo

import (
	"github.com/maxexq/isekei-shop-api/entities"
	"gorm.io/gorm"
)

type InventoryRepo interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, qty int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
