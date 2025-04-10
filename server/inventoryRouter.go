package server

import (
	_inventoryController "github.com/maxexq/isekei-shop-api/pkg/inventory/controller"
	_inventoryRepository "github.com/maxexq/isekei-shop-api/pkg/inventory/repo"
	_inventoryService "github.com/maxexq/isekei-shop-api/pkg/inventory/service"
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepoImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryService(
		inventoryRepository,
		itemShopRepository,
	)

	inventoryController := _inventoryController.NewInventoryController(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing) // /v1/inventory?playerID=1&serverID=1
}
