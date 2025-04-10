package server

import (
	_playerCoinController "github.com/maxexq/isekei-shop-api/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/maxexq/isekei-shop-api/pkg/playerCoin/repo"
	_playerCoinService "github.com/maxexq/isekei-shop-api/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepoImpl(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
	router.GET("", playerCoinController.Showing, m.PlayerAuthorizing)
}
