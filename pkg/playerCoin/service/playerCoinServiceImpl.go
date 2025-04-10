package service

import (
	_playerCoinRepo "github.com/maxexq/isekei-shop-api/pkg/playerCoin/repo"
)

type playerCoinServiceImpl struct {
	repo _playerCoinRepo.PlayerCoinRepo
}

func NewPlayerCoinServiceImpl(repo _playerCoinRepo.PlayerCoinRepo) PlayerCoinService {
	return &playerCoinServiceImpl{
		repo,
	}
}
