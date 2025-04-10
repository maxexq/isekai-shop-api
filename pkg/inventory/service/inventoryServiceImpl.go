package service

import (
	_inventoryRepo "github.com/maxexq/isekei-shop-api/pkg/inventory/repo"
	_itemShopRepo "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepo _inventoryRepo.InventoryRepo
	itemShopRepo  _itemShopRepo.ItemShopRepository
}

func NewInventoryService(inventoryRepo _inventoryRepo.InventoryRepo, itemShopRepo _itemShopRepo.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepo: inventoryRepo,
		itemShopRepo:  itemShopRepo,
	}
}
