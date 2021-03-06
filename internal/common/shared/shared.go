// Package shared is used to encapsulate items used by all other
// packages to prevent import cycles.
package shared

import "fmt"

// ClientID is an enum for client IDs
type ClientID int

const (
	// Team1 ID
	Team1 ClientID = iota
	// Team2 ID
	Team2
	// Team3 ID
	Team3
	// Team4 ID
	Team4
	// Team5 ID
	Team5
	// Team6 ID
	Team6
)

// TeamIDs contain sequential IDs of all teams
var TeamIDs = [...]ClientID{Team1, Team2, Team3, Team4, Team5, Team6}

func (c ClientID) String() string {
	clientIDStrings := [...]string{"Team1", "Team2", "Team3", "Team4", "Team5", "Team6"}
	if c >= 0 && int(c) < len(clientIDStrings) {
		return clientIDStrings[c]
	}
	return fmt.Sprintf("UNKNOWN ClientID '%v'", int(c))
}

// GoString implements GoStringer
func (c ClientID) GoString() string {
	return c.String()
}
