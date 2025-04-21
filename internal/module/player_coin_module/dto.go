package playercoinmodule

import "time"

type (
	PlayerCoin struct {
		ID        uint64    `json:"id"`
		PlayerID  string    `json:"playerID"`
		Amount    int64     `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	AddCoinRequest struct {
		PlayerID string `json:"playerID" validate:"required"`
		Amount   int64  `json:"amount" validate:"required,gt=0"`
	}

	ShowPlayerCoin struct {
		PlayerID string `json:"playerID" validate:"required"`
		Coin     int64  `json:"coin"`
	}
)
