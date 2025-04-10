package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/pkg/custom"
	_inventoryService "github.com/maxexq/isekei-shop-api/pkg/inventory/service"
	"github.com/maxexq/isekei-shop-api/pkg/validation"
)

type inventoryControllerImpl struct {
	service _inventoryService.InventoryService
	logger  echo.Logger
}

func NewInventoryController(
	service _inventoryService.InventoryService,
	logger echo.Logger,
) InventoryController {
	return &inventoryControllerImpl{
		service: service,
		logger:  logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)

	if err != nil {
		c.logger.Errorf("Player ID Getting error: %s", err.Error())

		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	inventoryListing, err := c.service.Listing(playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)
}
