package service

import (
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}
