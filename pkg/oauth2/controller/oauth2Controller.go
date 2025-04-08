package controller

import "github.com/labstack/echo/v4"

type OAuth2Controller interface {
	PlayerLogin(pctx echo.Context) error
	AdminLogin(pctx echo.Context) error
	PlayerLoginCallbak(pctx echo.Context) error
	AdminLoginCallback(pctx echo.Context) error
	Logout(pctx echo.Context) error
}
