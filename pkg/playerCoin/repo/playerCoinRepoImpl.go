package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
)

type playerCoinRepoImpl struct {
	db     *databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepoImpl(db *databases.Database, logger echo.Logger) PlayerCoinRepo {
	return &playerCoinRepoImpl{
		db:     db,
		logger: logger,
	}
}
