package itemshopapi

import "github.com/labstack/echo/v4"

type IItemShopHandler interface {
	GetItemShopList(c echo.Context) error
}
