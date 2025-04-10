package controller

import (
	_inventoryService "github.com/maxexq/isekei-shop-api/pkg/inventory/service"
)

type inventoryControllerImpl struct {
	service _inventoryService.InventoryService
}

func NewInventoryController(service _inventoryService.InventoryService) InventoryController {
	return &inventoryControllerImpl{
		service: service,
	}
}
