package itemshopapi

import (
	itemshopmodule "github.com/maxexq/isekei-shop-api/internal/module/itemshop_module"
)

type (
	APIGetItemShopListResponse struct {
		itemshopmodule.ListItemshopResponse
	}

	APIGetItemShopListFilter struct {
		itemshopmodule.ListItemshopRequestFilter
	}
)

type (
	APIPostBuyItemRequest struct {
		itemshopmodule.BuyItemRequest
	}

	APIPostSellItemRequest struct {
		itemshopmodule.SellItemRequest
	}
)
