package server

import (
	_itemShopController "github.com/maxexq/isekei-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/maxexq/isekei-shop-api/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopController(itemShopService)

	router.GET("", itemShopController.Listing)
}
