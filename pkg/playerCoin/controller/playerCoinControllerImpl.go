package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/pkg/custom"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
	_playerCoinService "github.com/maxexq/isekei-shop-api/pkg/playerCoin/service"
	"github.com/maxexq/isekei-shop-api/pkg/validation"
)

type PlayerCoinControllerImpl struct {
	service _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(service _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &PlayerCoinControllerImpl{
		service,
	}
}

func (c *PlayerCoinControllerImpl) CoinAdding(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)
	coinAddingReq.PlayerID = playerID

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoin, err := c.service.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *PlayerCoinControllerImpl) Showing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoinShowing := c.service.Showing(playerID)

	return pctx.JSON(http.StatusOK, playerCoinShowing)
}
