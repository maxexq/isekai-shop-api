package server

import (
	_itemManagingController "github.com/maxexq/isekei-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/maxexq/isekei-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/maxexq/isekei-shop-api/pkg/itemManaging/service"
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingService(
		itemManagingRepository,
		itemShopRepository,
	)
	itemManagingController := _itemManagingController.NewItemManagingController(itemManagingService)

	router.POST("", itemManagingController.Creating, m.AdminAuthorizing)
	router.PATCH("/:itemID", itemManagingController.Editing, m.AdminAuthorizing)
	router.DELETE("/:itemID", itemManagingController.Archiving, m.AdminAuthorizing)
}
