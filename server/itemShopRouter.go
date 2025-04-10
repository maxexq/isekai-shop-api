package server

import (
	_itemShopController "github.com/maxexq/isekei-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/maxexq/isekei-shop-api/pkg/itemShop/service"

	_inventoryRepository "github.com/maxexq/isekei-shop-api/pkg/inventory/repo"
	_playerCoinRepository "github.com/maxexq/isekei-shop-api/pkg/playerCoin/repo"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepoImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepoImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository, playerCoinRepository, inventoryRepository, s.app.Logger)
	itemShopController := _itemShopController.NewItemShopController(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)   // /v1/item-shop/buying?playerID=1&serverID=1&itemID=1&amount=1
	router.POST("/selling", itemShopController.Selling, m.PlayerAuthorizing) // /v1/item-shop/selling?playerID=1&serverID=1&itemID=1&amount=1
}
