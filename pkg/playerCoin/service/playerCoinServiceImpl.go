package service

import (
	"github.com/maxexq/isekei-shop-api/entities"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
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

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoin, err := s.repo.CoinAdding(nil, playerCoinEntity)
	if err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	playerCoinShowing, err := s.repo.Showing(playerID)

	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}

	return playerCoinShowing
}
