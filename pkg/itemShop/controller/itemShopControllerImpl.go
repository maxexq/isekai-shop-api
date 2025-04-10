package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/maxexq/isekei-shop-api/pkg/custom"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
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
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
