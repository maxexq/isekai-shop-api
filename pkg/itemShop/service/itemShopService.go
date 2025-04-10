package service

import (
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
}
