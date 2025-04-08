package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"
	_playerException "github.com/maxexq/isekei-shop-api/pkg/player/exception"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) PlayerRepository {
	return &playerRepositoryImpl{
		db,
		logger,
	}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Errorf("creating admin failed: %s", err.Error())

		return nil, &_playerException.PlayerNotFound{
			PlayerID: playerEntity.ID,
		}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Find player by ID failed: %s", err.Error())

		return nil, &_playerException.PlayerNotFound{
			PlayerID: playerID,
		}
	}

	return player, nil
}
