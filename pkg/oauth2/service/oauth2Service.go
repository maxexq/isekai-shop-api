package service

import (
	_adminModel "github.com/maxexq/isekei-shop-api/pkg/admin/model"
	_playerModel "github.com/maxexq/isekei-shop-api/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(AdminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}
