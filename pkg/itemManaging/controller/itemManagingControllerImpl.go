package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/pkg/custom"
	_itemManagingModel "github.com/maxexq/isekei-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/maxexq/isekei-shop-api/pkg/itemManaging/service"
	"github.com/maxexq/isekei-shop-api/pkg/validation"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingController(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{
		itemManagingService,
	}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	adminID, err := validation.AdminIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingControllerImpl) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemEditingReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemEditingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Editing(itemID, itemEditingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingControllerImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}

func (c *itemManagingControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")

	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}
