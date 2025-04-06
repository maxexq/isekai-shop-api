package controller

import (
	_itemShopService "github.com/maxexq/isekei-shop-api/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopController(
	itemShopService _itemShopService.ItemShopService,
) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
