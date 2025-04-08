package exception

import "fmt"

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return fmt.Sprintf("creating playerID: %s failed", e.PlayerID)
}
