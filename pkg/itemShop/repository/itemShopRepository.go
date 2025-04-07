package repository

import (
	"github.com/maxexq/isekei-shop-api/entities"
)

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
