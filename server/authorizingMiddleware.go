package server

import (
	"github.com/labstack/echo/v4"
	"github.com/maxexq/isekei-shop-api/config"
	_oauth2Controller "github.com/maxexq/isekei-shop-api/pkg/oauth2/controller"
)

type authorizingMiddleware struct {
	oauth2Controller _oauth2Controller.OAuth2Controller
	oauth2Conf       *config.OAuth2
	logger           echo.Logger
}

func (m *authorizingMiddleware) PlayerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauth2Controller.PlayerAuthorizing(pctx, next)
	}
}

func (m *authorizingMiddleware) AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauth2Controller.AdminAuthorizing(pctx, next)
	}
}
