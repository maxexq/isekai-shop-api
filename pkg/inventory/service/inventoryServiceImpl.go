package service

import (
	"github.com/maxexq/isekei-shop-api/entities"
	_inventoryRepo "github.com/maxexq/isekei-shop-api/pkg/inventory/repo"
	_itemShopRepo "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"

	_inventoryModel "github.com/maxexq/isekei-shop-api/pkg/inventory/model"
)

type inventoryServiceImpl struct {
	inventoryRepo _inventoryRepo.InventoryRepo
	itemShopRepo  _itemShopRepo.ItemShopRepository
}

func NewInventoryService(inventoryRepo _inventoryRepo.InventoryRepo, itemShopRepo _itemShopRepo.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepo: inventoryRepo,
		itemShopRepo:  itemShopRepo,
	}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventoryEntities, err := s.inventoryRepo.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventoryEntities)

	return s.buildInventoryListingResult(uniqueItemWithQuantityCounterList), nil
}

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(
	inventoryEntities []*entities.Inventory,
) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)
	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) buildInventoryListingResult(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []*_inventoryModel.Inventory {

	uniqueItemIDList := s.getItemID(uniqueItemWithQuantityCounterList)

	results := make([]*_inventoryModel.Inventory, 0)

	itemEntities, err := s.itemShopRepo.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})

	}

	return results
}

func (s *inventoryServiceImpl) getItemID(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
