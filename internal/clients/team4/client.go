// Package team4 contains code for team 4's client implementation
package team4

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

const id = shared.Team4

func init() {
	common.RegisterClient(id, &client{Client: common.NewClient(id)})
}

type client struct {
	common.Client
}

// StartOfTurnUpdate is updates the gamestate of the client at the start of each turn.
// The gameState is served by the server.
// OPTIONAL. Base should be able to handle it but feel free to implement your own.
func (c *client) StartOfTurnUpdate(gameState common.GameState) {
	c.Logf("Received game state update: %v", gameState)
	// TODO
}

// EndOfTurnActions executes and returns the actions done by the client that turn.
// OPTIONAL. Base should be able to handle it but feel free to implement your own.
func (c *client) EndOfTurnActions() []common.Action {
	c.Logf("EndOfTurnActions")
	return nil
}
