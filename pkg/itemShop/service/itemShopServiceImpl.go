package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/entities"
	_itemShopException "github.com/maxexq/isekei-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/maxexq/isekei-shop-api/pkg/itemShop/repository"
	_playerCoinModel "github.com/maxexq/isekei-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/maxexq/isekei-shop-api/pkg/playerCoin/repo"

	_inventoryRepository "github.com/maxexq/isekei-shop-api/pkg/inventory/repo"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepo
	inventoryRepository  _inventoryRepository.InventoryRepo
	logger               echo.Logger
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerCoinRepository _playerCoinRepository.PlayerCoinRepo,
	inventoryRepository _inventoryRepository.InventoryRepo,
	logger echo.Logger,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		logger,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)

	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page
	totalPage := s.totalPageCalculation(itemCounting, size)
	result := s.toItemResultResponse(itemList, page, totalPage)

	return result, nil
}

// 1. Find item by ID
// 2. Total price calculation
// 3. Check if player has enough coins
// 4. Deduct coins from player
// 5. Create purchase history
// 6. Return the purchased item and updated player coins
func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntities, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntities.ToItemModel(), buyingReq.Quantity)

	if err := s.checkPlayerCoin(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          itemEntities.ID,
		ItemName:        itemEntities.Name,
		ItemDescription: itemEntities.Description,
		ItemPrice:       itemEntities.Price,
		ItemPicture:     itemEntities.Picture,
		Quantity:        buyingReq.Quantity,
	})

	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Purchase history recorded: %s", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})

	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Player coin deducted: %s", playerCoin.Amount)

	inventoryEntity, err := s.inventoryRepository.Filling(
		tx, buyingReq.PlayerID,
		itemEntities.ID,
		int(buyingReq.Quantity),
	)
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Inventory filled: %s", inventoryEntity[0].ID)

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

// Selling handles the selling process of an item in the shop.
// 1. Find item by ID
// 2. Check if item is available for sale
// 3. Calculate the selling price
// 4. Add coins to player
// 5. Create a selling history
// 6. Return the updated player coins
// 7. Return the sold item and updated player coins
func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	return nil, nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)

	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) checkPlayerCoin(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Errorf("Player coin is not enough: %s", err)

		return &_itemShopException.CoinNotEnough{}
	}

	return nil
}
