package model

type (
	PlayerCoin struct {
		ID        uint64 `json:"id"`
		PlayerID  uint64 `json:"playerID"`
		Amount    int64  `json:"amount"`
		CreatedAt string `json:"createdAt"`
	}

	CoinAddingReq struct {
		PlayerID uint64 `json:"playerID" validate:"required"`
		Amount   int64  `json:"amount" validate:"required,gt=0"`
	}

	PlayerCoinShowingReq struct {
		PlayerID uint64 `json:"playerID" validate:"required"`
		Coin     int64  `json:"coin"`
	}
)
