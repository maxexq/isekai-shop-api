package service

import (
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item, error)
}
