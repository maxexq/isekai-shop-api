package service

import (
	_inventoryModel "github.com/maxexq/isekei-shop-api/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
