package repo

import (
	"github.com/maxexq/isekei-shop-api/entities"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepo interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
