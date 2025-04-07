package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/pkg/custom"
	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
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

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, (&_itemShopException.ItemListing{}).Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
