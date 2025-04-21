package itemshopmodule

import (
	playercoinmodule "github.com/maxexq/isekei-shop-api/internal/module/player_coin_module"
)

type IModule interface {
	ListItem(filter ListItemshopRequestFilter) (ListItemshopResponse, error)
	BuyItem(req BuyItemRequest) (playercoinmodule.PlayerCoin, error)
	SellItem(req SellItemRequest) (playercoinmodule.PlayerCoin, error)
}

type ModuleService struct{}
