package controller

import (
	_playerCoinService "github.com/maxexq/isekei-shop-api/pkg/playerCoin/service"
)

type PlayerCoinControllerImpl struct {
	service _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(service _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &PlayerCoinControllerImpl{
		service,
	}
}
