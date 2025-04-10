package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/databases"
	"github.com/maxexq/isekei-shop-api/entities"
	"gorm.io/gorm"

	_inventoryException "github.com/maxexq/isekei-shop-api/pkg/inventory/exception"
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

func (r *inventoryRepoImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntitiesResult := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntitiesResult = append(inventoryEntitiesResult, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(&inventoryEntitiesResult).Error; err != nil {
		r.logger.Errorf("filing inventory failed: %s", err.Error())

		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntitiesResult, nil
}

func (r *inventoryRepoImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, qty int) error {

	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities, err := r.findPlayerItemInInventoryByID(
		playerID,
		itemID,
		qty,
	)

	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {

		inventory.IsDeleted = true
		if err := conn.Model(&entities.Inventory{}).Where(
			"id = ?", inventory.ID,
		).Updates(inventory).Error; err != nil {
			r.logger.Errorf("removing player item in inventory failed: %s", err.Error())

			return &_inventoryException.PlayerItemRemoving{
				ItemID: itemID,
			}
		}

	}

	return nil
}

func (r *inventoryRepoImpl) findPlayerItemInInventoryByID(
	playerID string,
	itemID uint64,
	limit int,
) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?",
		playerID, itemID, false,
	).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("finding player item in inventory failed: %s", err.Error())
		return nil, &_inventoryException.PlayerItemRemoving{
			ItemID: itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepoImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	inventoryCount := int64(0)

	if err := r.db.Connect().Model(&entities.Inventory{}).Where(
		"player_id = ? and item_id = ? and is_deleted = ?",
		playerID, itemID, false,
	).Count(&inventoryCount).Error; err != nil {
		r.logger.Errorf("counting player item in inventory failed: %s", err.Error())

		return -1
	}

	return inventoryCount
}

func (r *inventoryRepoImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id = ? and is_deleted = ?", playerID, false).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("listing player inventory failed: %s", err.Error())

		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
		}
	}

	return inventoryEntities, nil
}
