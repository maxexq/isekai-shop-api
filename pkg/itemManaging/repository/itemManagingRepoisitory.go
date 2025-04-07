package repository

import (
	"github.com/maxexq/isekei-shop-api/entities"
	_itemManagingModel "github.com/maxexq/isekei-shop-api/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
}
