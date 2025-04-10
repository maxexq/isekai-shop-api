package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
)

type inventoryRepoImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepoImpl(db databases.Database, logger echo.Logger) InventoryRepo {
	return &inventoryRepoImpl{
		db:     db,
		logger: logger,
	}
}
