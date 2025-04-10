package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"

	_playerCoinException "github.com/maxexq/isekei-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
)

type playerCoinRepoImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepoImpl(db databases.Database, logger echo.Logger) PlayerCoinRepo {
	return &playerCoinRepoImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinRepoImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	playerCoin := new(entities.PlayerCoin)

	if err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Errorf("creating player coin failed: %s", err.Error())

		return nil, &_playerCoinException.CoinAdding{}
	}

	return playerCoin, nil
}

func (r *playerCoinRepoImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoinShowing := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(&entities.PlayerCoin{}).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group("player_id").Scan(playerCoinShowing).Error; err != nil {
		r.logger.Errorf("showing player coin failed: %s", err.Error())

		return nil, &_playerCoinException.PlayerCoinShowing{}
	}

	return playerCoinShowing, nil
}
