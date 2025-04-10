package service

import (
	"github.com/maxexq/isekei-shop-api/entities"
	_adminRepository "github.com/maxexq/isekei-shop-api/pkg/admin/repository"
	_playerRepository "github.com/maxexq/isekei-shop-api/pkg/player/repository"

	_adminModel "github.com/maxexq/isekei-shop-api/pkg/admin/model"
	_playerModel "github.com/maxexq/isekei-shop-api/pkg/player/model"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	if !s.IsThisGuyIsReallyPlayer((playerCreatingReq.ID)) {
		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Name:   playerCreatingReq.Name,
			Email:  playerCreatingReq.Email,
			Avatar: playerCreatingReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(AdminCreatingReq *_adminModel.AdminCreatingReq) error {
	if !s.IsThisGuyIsReallyAdmin((AdminCreatingReq.ID)) {
		adminEntity := &entities.Admin{
			ID:     AdminCreatingReq.ID,
			Name:   AdminCreatingReq.Name,
			Email:  AdminCreatingReq.Email,
			Avatar: AdminCreatingReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) PlayerAccountUpdating(AdminCreatingReq *_adminModel.AdminCreatingReq) error {
	return nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}

	return admin != nil
}
