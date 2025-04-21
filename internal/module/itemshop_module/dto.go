package itemshopmodule

import "github.com/maxexq/isekei-shop-api/pkg/web"

type (
	ItemShopItem struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Price       uint   `json:"price"`
	}

	ListItemshopRequestFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
		web.Paginate
	}

	ListItemshopResponse struct {
		Items    []*ItemShopItem `json:"items"`
		Paginate web.Pagination  `json:"paginate"`
	}
)

type (
	BuyItemRequest struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}

	SellItemRequest struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}
)
